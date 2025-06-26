package test

import (
	baseEntity "dzhgo/internal/model/entity"
	"dzhgo/utility/excel"
	"github.com/gogf/gf/contrib/config/apollo/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gcfg"
	_ "github.com/gzdzh-cn/dzhcore/contrib/drivers/mysql"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/smtp"
	"reflect"

	"context"
	logic "dzhgo/addons/crm/logic/sys"
	"dzhgo/addons/customer_pro/model/do"
	util2 "dzhgo/addons/customer_pro/utility/util"
	baseDao "dzhgo/internal/dao"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gzdzh-cn/dzhcore"
	"github.com/gzdzh-cn/dzhcore/utility/util"
	"github.com/shopspring/decimal"
	"os"
	"regexp"
	"sync"
	"time"

	"testing"

	customerDao "dzhgo/addons/customer_pro/dao"
	customerEntity "dzhgo/addons/customer_pro/model/entity"
	_ "dzhgo/internal/model"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/go-mail/mail"
	//"github.com/hraban/opus"
)

var (
	ctx     context.Context
	userMap = make(map[string]string)
)

func TestSql(t *testing.T) {
	_, err := customerDao.AddonsCustomerProKf.Ctx(ctx).Data(g.Map{"hasReceive": 0}).Update()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}
}

func TestExcel2(t *testing.T) {

	// 记录开始时间
	startTime := gtime.Now()
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()
	// 创建一个新的工作表
	sheetName := "Sheet1"
	sw, _ := f.NewStreamWriter(sheetName)
	styleID, err := f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "777777"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	// 使用 SetRow 方法设置从 A1 开始的行数据和样式
	if err := sw.SetRow("A1",
		[]interface{}{
			excelize.Cell{StyleID: styleID, Value: "Data"},
			[]excelize.RichTextRun{
				{Text: "Rich ", Font: &excelize.Font{Color: "2354E8"}},
				{Text: "Text", Font: &excelize.Font{Color: "E83723"}},
			},
		},
		excelize.RowOpts{Height: 45, Hidden: false}); err != nil {
		fmt.Println("设置行数据失败:", err)
		return
	}
	for rowID := 2; rowID <= 1152400; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, err := excelize.CoordinatesToCellName(1, rowID)
		if err != nil {
			fmt.Println(err)
			break
		}
		if err := sw.SetRow(cell, row); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("行数:", rowID)
		time.Sleep(100 * time.Microsecond)
	}
	if err := sw.Flush(); err != nil {
		fmt.Println(err)
		return
	}
	// 保存 Excel 文件
	if err := f.SaveAs("../public/excel/example.xlsx"); err != nil {
		fmt.Println("保存文件失败:", err)
	}
	// 计算执行时间
	executionTime := gtime.Now().Sub(startTime)
	g.Log().Infof(ctx, "数据已成功流式写入文件:%v，执行时间: %s", "/public/excel/example.xlsx", executionTime)
}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Item2 定义第二个结构体
type Item2 struct {
	Id      int    `json:"id"`
	Age     string `json:"age"`
	Address string `json:"address"`
	// 可以添加更多字段
}

// MergedItem 定义合并后的结构体
type MergedItem struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
	// 确保 MergedItem 包含 Item2 的所有字段
}

// mergeStructs 使用反射合并两个结构体
func mergeStructs(dst, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src)

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Type().Field(i)
		dstField := dstValue.FieldByName(srcField.Name)
		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(srcValue.Field(i))
		}
	}
}

// 每个 goroutine 写入的数据行数
const chunkSize = 1000

func writeDataChunk(f *excelize.File, sheetName string, startRow int, numRows int) {
	for i := 0; i < numRows; i++ {
		row := startRow + i
		// 逐行写入数据
		data := []string{
			fmt.Sprintf("姓名%d", row+1),
			fmt.Sprintf("%d", 20+row%10),
			fmt.Sprintf("城市%d", row%5),
		}
		for colIndex, value := range data {
			// 计算当前单元格的位置
			cell, err := excelize.CoordinatesToCellName(colIndex+1, row+2)
			if err != nil {
				fmt.Println("生成单元格名称时出错:", err)
				return
			}
			// 设置单元格的值
			f.SetCellValue(sheetName, cell, value)
		}
	}
}

func TestExcelLot(t *testing.T) {
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()
	// 获取默认工作表的名称
	sheetName := f.GetSheetName(f.GetActiveSheetIndex())

	// 写入表头
	headers := []string{"姓名", "年龄", "城市"}
	for colIndex, header := range headers {
		// 将列索引和行索引转换为 Excel 单元格名称
		cell, err := excelize.CoordinatesToCellName(colIndex+1, 1)
		if err != nil {
			fmt.Println("生成单元格名称时出错:", err)
			return
		}
		// 设置单元格的值为表头
		f.SetCellValue(sheetName, cell, header)
	}

	// 总数据行数
	totalRows := 350000
	var wg sync.WaitGroup

	// 分块并发写入数据
	for i := 0; i < totalRows; i += chunkSize {
		wg.Add(1)
		end := i + chunkSize
		if end > totalRows {
			end = totalRows
		}
		go func(start int, num int) {
			defer wg.Done()
			writeDataChunk(f, sheetName, start, num)
		}(i, end-i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 最后保存文件
	if err := f.SaveAs("../public/output_large.xlsx"); err != nil {
		fmt.Println("保存 Excel 文件时出错:", err)
	}
	fmt.Println("数据已成功流式写入 output_large.xlsx")
}

func TestMerge(t *testing.T) {
	var itemAttr1 []*Item
	var itemAttr2 []*Item2

	data := `[{"id":1,"name":"3"},{"id":2,"name":"4"},{"id":3,"name":"5"}]`
	data2 := `[{"id":1,"age":"30","address":"addr1"},{"id":2,"age":"40","address":"addr2"},{"id":3,"age":"50","address":"addr3"}]`

	err := gconv.Struct(data, &itemAttr1)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	err = gconv.Struct(data2, &itemAttr2)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	attrMap := make(map[int]*Item2)
	for _, row := range itemAttr2 {
		attrMap[row.Id] = row
	}

	//var mergedItems []MergedItem
	//for _, row := range *itemAttr1 {
	//	if item2, ok := attrMap[row.Id]; ok {
	//		mergedItem := MergedItem{
	//			Id:   row.Id,
	//			Name: row.Name,
	//		}
	//		mergeStructs(&mergedItem, item2)
	//		mergedItems = append(mergedItems, mergedItem)
	//	}
	//}
	//
	//g.Dump(mergedItems)

	var mergedItems2 []*MergedItem
	for _, row := range itemAttr1 {
		if item2, ok := attrMap[row.Id]; ok {
			myMap1 := gmap.NewStrAnyMapFrom(gconv.Map(row))
			myMap2 := gmap.NewStrAnyMapFrom(gconv.Map(item2))
			myMap1.Merge(myMap2)

			var mergedItem *MergedItem
			err = gconv.Struct(myMap1.Map(), &mergedItem)
			if err != nil {
				g.Log().Error(ctx, err)
			}
			mergedItems2 = append(mergedItems2, mergedItem)
		}
	}
	g.Dump(mergedItems2)

}

//func tranOpus() {
//	// 打开 Opus 文件
//	inputFile, err := os.Open("input.opus")
//	if err != nil {
//		fmt.Println("Error opening input file:", err)
//		return
//	}
//	defer inputFile.Close()
//
//	// 创建 WAV 文件
//	outputFile, err := os.Create("output.wav")
//	if err != nil {
//		fmt.Println("Error creating output file:", err)
//		return
//	}
//	defer outputFile.Close()
//
//	// 初始化 Opus 解码器
//	decoder, err := opus.NewDecoder(48000, 2)
//	if err != nil {
//		fmt.Println("Error initializing Opus decoder:", err)
//		return
//	}
//
//	// 初始化 WAV 编码器
//	enc := wav.NewEncoder(outputFile, 48000, 16, 2, 1)
//
//	// 缓冲区
//	frameSize := 960 // 20ms frame at 48kHz
//	buffer := make([]byte, 4096)
//	pcm := make([]int16, frameSize*2)
//
//	for {
//		n, err := inputFile.Read(buffer)
//		if err != nil && err != io.EOF {
//			fmt.Println("Error reading input file:", err)
//			return
//		}
//		if n == 0 {
//			break
//		}
//
//		// 解码 Opus 数据
//		numSamples, err := decoder.Decode(buffer[:n], pcm)
//		if err != nil {
//			fmt.Println("Error decoding Opus data:", err)
//			return
//		}
//
//		// 转换 []int16 到 []int
//		intBuffer := make([]int, numSamples*2)
//		for i := 0; i < numSamples*2; i++ {
//			intBuffer[i] = int(pcm[i])
//		}
//
//		// 写入 WAV 文件
//		audioBuffer := &audio.IntBuffer{Data: intBuffer, Format: &audio.Format{SampleRate: 48000, NumChannels: 2}}
//		if err := enc.Write(audioBuffer); err != nil {
//			fmt.Println("Error writing to WAV file:", err)
//			return
//		}
//	}
//
//	// 关闭 WAV 编码器
//	if err := enc.Close(); err != nil {
//		fmt.Println("Error closing WAV encoder:", err)
//		return
//	}
//
//	fmt.Println("Conversion complete!")
//}

func findClues() {
	userIds, _ := customerDao.AddonsCustomerProKf.Ctx(ctx).Where("groupId", "1831215215165837312").Fields("userId").Array()
	var sql string
	for i, userId := range userIds {

		if i == len(userIds)-1 {
			sql += fmt.Sprintf("FIND_IN_SET(%v, services_ids) ", userId)
		} else {
			sql += fmt.Sprintf("FIND_IN_SET(%v, services_ids) or ", userId)
		}
	}
	fmt.Println(sql)
}

func doExcel2() {

	type Clues struct {
		ID           int       `json:"id"`            // 记录ID
		SerialID     string    `json:"serialId"`      // 序列号或唯一标识符
		Name         string    `json:"name"`          // 姓名
		Mobile       string    `json:"mobile"`        // 手机号码
		WeChat       string    `json:"wechat"`        // 微信账号
		Keywords     string    `json:"keywords"`      // 关键词
		FollowUpType string    `json:"followupType"`  // 跟进类型
		Status       string    `json:"status"`        // 状态
		AccountId    string    `json:"account_id"`    //
		ServicesIds  string    `json:"services_ids"`  //
		SourceFrom   string    `json:"source_from"`   // 来源
		GuestIPInfo  string    `json:"guest_ip_info"` // 访客IP信息
		CreateTime   time.Time `json:"createTime"`    // 创建时间
	}

	var (
		accountMap  = g.MapStrAny{} //id为下标的账号数集
		userNameMap = g.MapStrStr{} //id为下标的会员数集
	)

	//id为下标的账号数集
	{
		mApi := customerDao.AddonsCustomerProApi.Ctx(ctx)
		mApi = mApi.Fields("id", "name")
		apiList, err := mApi.All()
		if err != nil {
			g.Log().Error(ctx, err.Error())
			//return err
		}
		for _, v := range apiList {
			accountMap[gconv.String(v["id"])] = gconv.String(v["name"])
		}
	}
	//id为下标的会员数集
	{
		var userList []*baseEntity.BaseSysUser
		err := baseDao.BaseSysUser.Ctx(ctx).Fields("id", "name").Scan(&userList)
		if err != nil {
			g.Log().Error(ctx, err.Error())
		}
		for _, user := range userList {
			userNameMap[user.Id] = user.Name
		}
	}

	ids := []string{
		"1853720097529532416",
		"1853720032836587520",
	}
	db := customerDao.AddonsCustomerProClues.Ctx(ctx)
	allClues, _ := db.WhereIn("id", ids).Fields("id,serialId,name,mobile,wechat,keywords,followupType,status,account_id,services_ids,source_from,guest_ip_info,createTime").All()
	var (
		sheetList [][]interface{}
		clues     *Clues
		mapData   []*excel.ExcelSheet
	)

	for _, record := range allClues {
		gconv.Struct(record.Map(), &clues)
		var list []interface{}
		list = append(list, clues.SerialID)
		list = append(list, clues.Name)
		list = append(list, clues.Mobile)
		list = append(list, clues.WeChat)
		list = append(list, clues.Keywords)
		list = append(list, util2.FollowUpType(gconv.Int(clues.FollowUpType)))
		list = append(list, util2.Status(gconv.Int(clues.Status)))
		list = append(list, accountMap[clues.AccountId])
		list = append(list, util2.ServicesIds(clues.ServicesIds, userNameMap))
		list = append(list, util2.SourceFrom(gconv.Int(clues.SourceFrom)))
		list = append(list, clues.GuestIPInfo)
		list = append(list, clues.CreateTime)
		sheetList = append(sheetList, list)
	}
	mapData = append(mapData, &excel.ExcelSheet{
		SheetName: "Sheet1",
		SheetHead: g.Array{
			"序号", "姓名", "手机号", "微信号", "关键词", "跟进状态", "成交状态", "账户", "分配过的客服",
			"来源", "IP归属地", "创建时间",
		},
		SheetList: sheetList,
	})
	e := excel.NewExcel()
	e.SetExcelSheetData(mapData, false).SaveExcel("./", "test.xlsx")

}

func doExcel() {
	headList := g.Array{"姓名", "年龄", "性别", "身高"}
	dataList := []g.Array{{"李", "12", "男", "170"}, {"任", "11", "女", "160"}}

	headList2 := g.Array{"姓名", "年龄", "性别"}
	dataList2 := []g.Array{{"李", "12", "男"}, {"任", "11", "女"}}
	e := excel.NewExcel()
	e.SetSheet("Sheet1").SetCellHead(headList).SetCellRow(dataList, 1)
	e.SetSheet("Sheet2").SetCellHead(headList2).SetCellRow(dataList2, 2)
	e.SaveExcel("./", "test.xlsx")
}

func ExcelTo() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个工作表
	_, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)

	// 设置工作簿的默认工作表
	//f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func gosmtp() {

	os.Setenv("GODEBUG", "tlsrsakex=1")
	username := "lilie2010@sina.com"
	password := "5d3b28ce4fe1f640"
	smtpHost := "smtp.sina.com"
	auth := smtp.PlainAuth("", username, password, smtpHost)
	from := "lilie2010@sina.com"
	to := []string{"liliewin@163.com"}
	message := []byte("To: liliewin@163.com\r\n" +
		"From: lilie2010@sina.com\r\n" +
		"\r\n" +
		"Subject: Why aren't you using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here's the space for your great sales pitch\r\n")
	smtpUrl := smtpHost + ":25"
	err := smtp.SendMail(smtpUrl, auth, from, to, message)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
}

func gomail() {

	os.Setenv("GODEBUG", "tlsrsakex=1")

	var (
		Addr     = "smtp.sina.com"
		AuthName = "from@sina.com"
		AuthPwd  = "123456"
	)
	m := mail.NewMessage()
	m.SetHeader("From", AuthName)
	m.SetHeader("To", "gzdazhihui@163.com")
	m.SetAddressHeader("Cc", "gzdazhihui@163.com", "zheng")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")
	d := mail.NewDialer(Addr, 25, AuthName, AuthPwd)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func sentEmail() {

	os.Setenv("GODEBUG", "tlsrsakex=1")
	var (
		Addr     = "smtp.sina.com"
		Host     = Addr + ":25"
		AuthName = "lilie2010@sina.com"
		AuthPwd  = "5d3b28ce4fe1f640"
		from     = "lilie2010@sina.com"
		To       = []string{"liliewin@163.com"}
		msg      = []byte("To: liliewin@163.com\r\n" +
			"From: lilie2010@sina.com\r\n" +
			"\r\n" +
			"Subject: Why aren't you using Mailtrap yet?\r\n" +
			"\r\n" +
			"Here's the space for your great sales pitch\r\n")
	)

	auth := smtp.PlainAuth("", AuthName, AuthPwd, Addr)
	if err := smtp.SendMail(Host, auth, from, To, msg); err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}
}

func exampelAdpter() {
	c, err := gcfg.New()
	if err != nil {
		panic(err)
	}
	adapter := c.GetAdapter()
	c.SetAdapter(adapter)
	fmt.Println(adapter)
}

func adapterTest() {

	// 加载配置
	//cfg := g.Cfg(guid.S())
	//configFile := flag.String("f", "/manifest/config/config-dev.yaml", "The config file path")
	//flag.Parse()
	//
	//GetAdapter, ok := cfg.GetAdapter().(*gcfg.AdapterFile)
	//if !ok {
	//	g.Log().Fatalf(ctx, "Failed to assert configuration adapter type")
	//}
	//GetAdapter.SetFileName(*configFile)

	var (
		appId   = "SampleApp"
		cluster = "default"
		ip      = "http://localhost:8080"
	)

	adapter, err := apollo.New(ctx, apollo.Config{
		AppID:   appId,
		IP:      ip,
		Cluster: cluster,
	})
	if err != nil {
		g.Log().Fatalf(ctx, `%+v`, err)
	}
	g.Cfg().SetAdapter(adapter)

	fmt.Println("----------server:", g.Cfg().MustGet(ctx, "server"))

}

func tranStruct() {
	type SaveHistoryListReg struct {
		g.Meta        `path:"/path" tags:"test" method:"post" summary:"test"`
		Authorization string `in:"header" dc:"token"  `
		History       string
	}

	reqMap := map[string]interface{}{
		"history":       "123",
		"Authorization": "auth",
	}

	var req *SaveHistoryListReg
	if err := gconv.Struct(reqMap, &req); err != nil {
		panic(err)
	}
	g.Dump(reqMap, req)

}

func payOrder() {
	url := "https://mz.phpwc.com/submit.php?pid=1053&type=wxpay&out_trade_no=20240806151343349&notify_url=https://xx.xx.xx/notify_url.php&return_url=https://xx.xx.xx/return_url.php&name=VIP&money=1.00&sign=7gmh5viJT3UyGotIShyLPTYab7JkFGvA&sign_type=MD5"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func sentMail() {
	_, err := logic.NewsOrderService().ServiceSentMail(ctx, true)
	if err != nil {
		return
	}
}

func humanizeFun() {
	bytes := uint64(9133600000)                                  // 示例字节数
	fmt.Println("Human-readable format:", humanize.Bytes(bytes)) // 输出易读格式
}

func getFileSize() {

	g.Log().Debugf(ctx, "get file size:%v", gfile.Size("./data/logs/2.log"))
}

// JsonOutputsForLogger is for JSON marshaling in sequence.
type JsonOutputsForLogger struct {
	Time     string `json:"time"`
	Level    string `json:"level"`
	Content  string `json:"content"`
	FuncName string `json:"funcName"`
}

// LoggingJsonHandler is a example handler for logging JSON format content.
var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	jsonForLogger := JsonOutputsForLogger{
		Time:     in.TimeFormat,
		Level:    gstr.Trim(in.LevelFormat, "[]"),
		Content:  gstr.Trim(in.ValuesContent()),
		FuncName: "方法名称",
	}
	jsonBytes, err := json.Marshal(jsonForLogger)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return
	}

	in.Buffer.Write(jsonBytes)
	in.Buffer.WriteString("\n")
	in.Next(ctx)
}

var LogHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {

	//r := g.RequestFromCtx(ctx)

	in.Buffer.WriteString(in.ValuesContent())
	in.Buffer.WriteString("\n")
	in.Buffer.WriteString("[ info ] [ 接口 ]")
	//in.Buffer.WriteString(r.URL.Path)

	in.Next(ctx)
}

type MyWriter struct {
	logger *glog.Logger
}

//func (w *MyWriter) Write(p []byte) (n int, err error) {
//
//	return w.g.Log().Write(p)
//}

func ctxDiff() {
	ctx := gctx.New()
	ctx2 := context.Background()
	ctx3 := gctx.GetInitCtx()

	g.Dump(ctx, ctx2, ctx3)
}

func serialIdChange() {
	var cluesList []*do.AddonsCustomerProClues
	err := customerDao.AddonsCustomerProClues.Ctx(ctx).Where("id > ", 67114).Scan(&cluesList)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	num := 67114
	for _, clue := range cluesList {
		num++
		clue.SerialId = num
	}

	_, err = customerDao.AddonsCustomerProClues.Ctx(ctx).Data(cluesList).Save()
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func importData() {
	path := "../public/2024-09-07.log"
	//re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}.*(talk_info contentData:|customer contentData:).*$`)
	// 匹配以时间开头并包含 "talk_info contentData:" 或 "customer contentData:" 的行，并提取时间和冒号后的内容
	re := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}).*(talk_info|customer) contentData:\s*(.*)$`)
	i := 0
	err := gfile.ReadLines(path, func(text string) error {

		//提取 2024-09-07 10:34:39 到 2024-09-07 13:22:31
		// 定义时间范围
		startTime, _ := time.Parse("2006-01-02 15:04:05", "2024-09-07 10:34:39")
		endTime, _ := time.Parse("2006-01-02 15:04:05", "2024-09-07 14:00:00")
		matches := re.FindStringSubmatch(text)
		if len(matches) > 0 {
			// 解析匹配到的时间
			logTime, _ := time.Parse("2006-01-02 15:04:05.000", matches[1])

			// 检查时间是否在指定范围内
			if logTime.After(startTime) && logTime.Before(endTime) {
				i++
				fmt.Printf("i:%v, Time: %s, Type: %s, Content: %s\n", i, matches[1], matches[2], matches[3])
				restoreInsertData := &do.AddonsCustomerProCluesRestore{
					Id:         dzhcore.NodeSnowflake.Generate().String(),
					Type:       matches[2],
					Remark:     matches[3],
					CreateTime: gtime.NewFromStr(matches[1]),
					UpdateTime: gtime.NewFromStr(matches[1]),
				}
				_, err := customerDao.AddonsCustomerProCluesRestore.Ctx(ctx).Data(restoreInsertData).OmitEmpty().Insert()
				if err != nil {
					g.Log().Error(ctx, err.Error())
				}
			}
		}

		return nil
	})
	if err != nil {
		return
	}

}

func gstrSplit(ctx context.Context) {

	type CommonError struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	//SendErr := "get access_token error : errcode=40164 , errormsg=invalid ip 113.101.177.118 ipv6 ::ffff:113.101.177.118, not in whitelist rid: 66d9be16-212ff30f-0e99c2f0"
	//SendErr := " Error , errcode=47003 , errmsg=argument invalid! data.phone_number6.value invalid rid: 66d9c752-38b9ae8f-093cbc59 \n"

	jsonData := []byte(`{"errcode":200,"errmsg":"valid"}`)

	var errmsg CommonError
	if err := json.Unmarshal(jsonData, &errmsg); err != nil {
		g.Log().Error(ctx, err)
	}

	if errmsg.ErrCode != 0 {
		fmt.Println(fmt.Errorf("Error , errcode=%d , errmsg=%s", errmsg.ErrCode, errmsg.ErrMsg))
	}

}

// 事务测试
func transaction(ctx context.Context) {
	err := g.DB().Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		err := sqlTransaction2(ctx) //传递带事务的ctx
		if err != nil {
			return err
		}
		return gerror.New("出错了")
	})
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}

}

func gdbSql(ctx context.Context) {
	err := sqlTransaction2(ctx) //传递不带事务的ctx
	if err != nil {

		return
	}
}

func sqlTransaction2(ctx context.Context) error {

	//如果接收到带事务的ctx就会启动事务
	_, err := customerDao.AddonsCustomerProApi.Ctx(ctx).Data(g.Map{"id": 1, "name": "测试"}).Insert()
	if err != nil {
		panic(err)
	}

	return err
}

func sqlTransaction(ctx context.Context, db interface{}) error {

	var tx gdb.DB
	if v, ok := db.(gdb.TX); ok {
		tx = v.GetDB()

	} else {
		tx = g.DB()
	}

	//	其他逻辑
	_, err := tx.Ctx(ctx).Model(customerDao.AddonsCustomerProApi.Table()).Data(g.Map{"id": 1, "name": "测试"}).Insert()
	if err != nil {
		panic(err)
	}

	return err

}

func jsonTran() {
	type Admin struct {
		IsRefresh       bool     `json:"isRefresh"`
		RoleIds         []string `json:"roleIds"`
		Username        string   `json:"username"`
		UserId          string   `json:"userId"`
		PasswordVersion *int32   `json:"passwordVersion"`
	}

	var admin *Admin
	jsonStr := `{"isRefresh":false,"roleIds":[1],"username":"admin","userId":1152921504606846975,"passwordVersion":1}`

	err := gjson.New(jsonStr).Scan(&admin)
	if err != nil {
		fmt.Println(err)
	}

	g.Dump(admin)

}

type RegisterError struct {
	Msg  string
	Code int
}

func (e *RegisterError) Error() string {
	return fmt.Sprintf("错误消息:%d,%s", e.Code, e.Msg)
}

func (e *RegisterError) Error2() string {
	return fmt.Sprintf("错误消息2:%d,%s", e.Code, e.Msg)
}

func register(username string) error {
	if username != "admin" {
		return &RegisterError{
			Code: -1,
			Msg:  "账号不对",
		}
	}

	return nil
}

func getEnv() {

	// 获取 GOHOSTARCH 环境变量
	goHostArch := os.Getenv("GOHOSTARCH")
	fmt.Println("GOHOSTARCH:", goHostArch)

	// 获取 GCCGO 环境变量
	gccGo := os.Getenv("GCCGO")
	fmt.Println("GCCGO:", gccGo)

	// 获取 AR 环境变量
	ar := os.Getenv("AR")
	fmt.Println("AR:", ar)

	// 获取 CC 环境变量
	cc := os.Getenv("CC")
	fmt.Println("CC:", cc)

	// 获取 CXX 环境变量
	cxx := os.Getenv("CXX")
	fmt.Println("CXX:", cxx)

	// 获取 CGO_ENABLE 环境变量
	cgoEnable := os.Getenv("CGO_ENABLE")
	fmt.Println("CGO_ENABLE:", cgoEnable)

}

// 迁移
func migrateData() error {

	var wg sync.WaitGroup
	wg.Add(1)
	// 使用一个新的背景上下文
	newCtx := gctx.New()
	go func(ctx context.Context) {

		defer wg.Done()
		err := g.DB().Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

			//清空数据表
			tables := "base_sys_user,base_sys_role,base_sys_user_role,base_sys_role_menu,addons_customer_pro_clues,addons_customer_pro_clues_record,addons_customer_pro_clues_track,addons_customer_pro_kf,addons_customer_pro_majors,addons_customer_pro_order,addons_customer_pro_project,addons_customer_pro_project_group,addons_customer_pro_readdegree,addons_customer_pro_readtypes,addons_customer_pro_school,addons_customer_pro_wx_user,addons_customer_pro_chatContent"
			err = clearData(ctx, tx, tables)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}

			////会员和角色，客服列表
			//userProjectMap, userProjectGroupMap, err := roleAndUserMove(tx)
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return err
			//}
			//
			////项目和客服组，需要会员表数据
			//err = projectAndGroup(tx, userProjectMap, userProjectGroupMap)
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return err
			//}

			////学院数据 专业数据 报读类型 报读层次
			//err = schoolAndMajorsAndDegreeAndTypes(tx)
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return err
			//}
			//
			////线索
			//err = cluesAndService(tx)
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return err
			//}
			//
			////录入线索跟进
			//err = cluesRecord(tx)
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return err
			//}
			//
			////录入线索轨迹
			//err = cluesTracks(tx)
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return err
			//}

			// 微信用户
			err = wxUser(tx)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}

			return nil
		})
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}(newCtx)

	wg.Wait()

	return nil

}

// 清空数据表数据
func clearData(ctx context.Context, tx gdb.TX, tables string) error {

	tableSlice := gstr.Split(tables, ",")
	for _, table := range tableSlice {
		if table == "base_sys_role" || table == "base_sys_user" || table == "base_sys_user_role" {
			if table == "base_sys_role" {
				_, err := tx.Model(table).WhereNot("id", "1").Unscoped().Delete()
				if err != nil {
					g.Log().Error(ctx, err)
					return err
				}
			}
			if table == "base_sys_user" {
				_, err := tx.Model(table).Where("id NOT IN (?)", g.SliceStr{"1152921504606846975", "1152921504606846976"}).Unscoped().Delete()
				if err != nil {
					g.Log().Error(ctx, err)
					return err
				}
			}
			if table == "base_sys_user_role" {
				_, err := tx.Model(table).Where("id NOT IN (?)", g.SliceStr{"1152921504606846977", "1152921504606846978"}).Unscoped().Delete()
				if err != nil {
					g.Log().Error(ctx, err)
					return err
				}
			}

		} else {

			_, err := tx.Exec(`TRUNCATE TABLE ` + table)
			//_, err := tx.Model(table).Where("id IS NOT NULL").Unscoped().Delete()
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		}

		g.Log().Warningf(ctx, "清空：%v", table)
	}

	return nil
}

// 线索迁移
func cluesAndService(tx gdb.TX) error {

	var (
		cluesSlice    []*customerEntity.OiClues
		dzhCluesSlice []*g.Map
		cluesMap      = make(map[string]*customerEntity.OiClues)
	)

	whereData := g.Map{
		//"status":             1,
		"deletetime IS NULL": "",
	}

	//读取线索表
	err := customerDao.OiClues.Ctx(ctx).Where(whereData).Scan(&cluesSlice)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return err
	}

	for _, v := range cluesSlice {

		cluesMap[gconv.String(v.Id)] = v

		chatContentVersion := 1
		if v.ChatContent == "" {
			chatContentVersion = 0
		}
		status := 0
		if v.Status == 2 {
			status = 1
		}

		//当前客服id
		var servicesId string
		if v.ServicesIds != "" {
			str := gstr.Split(v.ServicesIds, ",")
			servicesId = str[len(str)-1]
		}

		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}

		dzhCluesSlice = append(dzhCluesSlice, &g.Map{
			"id":                 gconv.String(v.Id),
			"name":               v.Name,
			"account_id":         gconv.String(v.AccountId),
			"address":            v.Address,
			"allot_time":         gtime.New(v.AllotTime),
			"chat_content":       v.ChatContent,
			"chatContentVersion": chatContentVersion,
			"created_id":         gconv.String(v.CreatedId),
			"created_name":       userMap[gconv.String(v.CreatedId)],
			"degree_id":          gconv.String(v.DegreeId),
			"education":          gconv.String(v.Education),
			"emergency_mobile":   v.EmergencyMobile,
			"followup_type":      v.FollowupType,
			"gender":             v.Gender,
			"graduated_school":   gconv.String(v.GraduatedSchool),
			"guest_id":           gconv.String(v.GuestId),
			"keywords":           v.Keywords,
			"last_followup_time": gtime.New(v.LastFollowupTime),
			"level":              v.Level,
			"majors_id":          gconv.String(v.MajorsId),
			"majors_type":        gconv.String(v.MajorsType),
			"mobile":             v.Mobile,
			"ocean_time":         gtime.New(v.OceanTime),
			"project_id":         gconv.String(v.ProjectId),
			"school_id":          gconv.String(v.SchoolId),
			"services_ids":       v.ServicesIds,
			"services_id":        servicesId,
			"source_from":        v.SourceFrom,
			"status":             status,
			"household_address":  v.HouseholdAddress,
			"createTime":         createTime,
			"updateTime":         updateTime,
		})
	}

	// 录入线索表
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProClues.Table(), dzhCluesSlice, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	//录入线索订单
	err = cluesOrder(tx, cluesMap)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}

// 线索跟进
func cluesRecord(tx gdb.TX) error {

	var (
		oiRecordSlice  []*do.OiCluesRecords
		dzhRecordSlice []*do.AddonsCustomerProCluesRecord
	)
	whereData := g.Map{
		"deletetime IS NULL": "",
	}
	err := customerDao.OiCluesRecords.Ctx(ctx).Where(whereData).Scan(&oiRecordSlice)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	nextFollowupTime := gtime.Now()
	for _, v := range oiRecordSlice {
		if gconv.Int(v.NextFollowupTime) != 0 {
			nextFollowupTime = gtime.New(gconv.Int(v.NextFollowupTime))
		}
		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}

		dzhRecordSlice = append(dzhRecordSlice, &do.AddonsCustomerProCluesRecord{
			Id:               v.Id,
			CluesId:          v.CluesId,
			CreatedId:        v.UserId,
			FollowupType:     v.FollowupType,
			NextFollowupTime: nextFollowupTime,
			IsNoticed:        v.IsNoticed,
			Remark:           v.Desc,
			CreateTime:       createTime,
			UpdateTime:       updateTime,
		})
	}

	//写入跟进表
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProCluesRecord.Table(), dzhRecordSlice, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil

}

// 线索订单
func cluesOrder(tx gdb.TX, cluesMap map[string]*customerEntity.OiClues) error {

	var (
		oiOrder  []*do.OiCluesOrder
		dzhOrder []*do.AddonsCustomerProOrder
	)

	whereData := g.Map{
		"deletetime IS NULL": "",
	}
	err := customerDao.OiCluesOrder.Ctx(ctx).Where(whereData).Scan(&oiOrder)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	if oiOrder == nil {
		err = gerror.New("没有数据")
		g.Log().Error(ctx, err)
		return err
	}

	for _, v := range oiOrder {

		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}

		graduatedDate := gtime.Now()
		isNil := util.GetValueOrDefault(v.GraduatedDate)
		if isNil != false {
			graduatedDate = gtime.New(gconv.Int(v.GraduatedDate))
		}

		cluesId := gconv.String(v.CluesId)
		var (
			majorsId   int
			majorsType int
			projectId  int
			schoolId   int
		)

		if _, exists := cluesMap[cluesId]; exists {
			majorsId = cluesMap[cluesId].MajorsId
		}

		if _, exists := cluesMap[cluesId]; exists {
			majorsType = cluesMap[cluesId].MajorsType
		}

		if _, exists := cluesMap[cluesId]; exists {
			projectId = cluesMap[cluesId].ProjectId
		}

		if _, exists := cluesMap[cluesId]; exists {
			schoolId = cluesMap[cluesId].SchoolId
		}

		dzhOrder = append(dzhOrder, &do.AddonsCustomerProOrder{
			Id:               v.Id,
			CluesId:          v.CluesId,
			CreatedId:        "0",
			Address:          v.Address,
			AuditNote:        v.AuditNote,
			AuditStatus:      v.AuditStatus,
			Birthday:         nil,
			Education:        v.Education,
			EmergencyContact: v.EmergencyContact,
			EmergencyMobile:  v.EmergencyMobile,
			Freshman:         v.Freshman,
			Gender:           v.Gender,
			GraduatedDate:    graduatedDate,
			GraduatedSchool:  v.GraduatedSchool,
			HouseholdAddress: v.HouseholdAddress,
			HouseholdType:    v.HouseholdType,
			IdcardNumber:     v.IdcardNumber,
			MajorsId:         majorsId,
			MajorsType:       majorsType,
			ProjectId:        projectId,
			SchoolId:         schoolId,
			Mobile:           v.Mobile,
			Name:             v.Name,
			Nation:           v.Nation,
			NativePlace:      v.NativePlace,
			PoliticsStatus:   v.PoliticsStatus,
			Receiver:         v.Receiver,
			SchoolPayment:    v.SchoolPayment,
			Serial:           v.Serial,
			ServicesId:       v.ServicesId,
			Status:           v.Status,
			TeamsPayment:     v.TeamsPayment,
			Voucher:          v.Voucher,
			Remark:           v.Desc,
			CreateTime:       createTime,
			UpdateTime:       updateTime,
		})
	}

	//g.Dump(dzhOrder)

	//写入线索
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProOrder.Table(), dzhOrder, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}

// 录入线索轨迹
func cluesTracks(tx gdb.TX) error {

	var (
		oiTracks  []*do.OiCluesTracks
		dzhTracks []*do.AddonsCustomerProCluesTrack
	)

	whereData := g.Map{
		"deletetime IS NULL": "",
	}
	err := customerDao.OiCluesTracks.Ctx(ctx).Where(whereData).Scan(&oiTracks)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	for _, v := range oiTracks {
		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}
		dzhTracks = append(dzhTracks, &do.AddonsCustomerProCluesTrack{
			Id:         v.Id,
			CluesId:    v.CluesId,
			CreatedId:  v.UserId,
			Type:       v.Type,
			Remark:     v.Note,
			CreateTime: createTime,
			UpdateTime: updateTime,
		})
	}

	//写入线索轨迹
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProCluesTrack.Table(), dzhTracks, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}

// 学院数据
func schoolAndMajorsAndDegreeAndTypes(tx gdb.TX) error {

	var (
		oiSchool  []*do.OiSchool
		dzhSchool []*do.AddonsCustomerProSchool
	)

	whereData := g.Map{
		"deletetime IS NULL": "",
	}

	//读取学校信息
	err := customerDao.OiSchool.Ctx(ctx).Where(whereData).Scan(&oiSchool)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	logo := "/dzhimg/assets/adminking.png"

	for _, v := range oiSchool {

		if v.Logo != "" {
			logo = gstr.Replace(gconv.String(v.Logo), "http://ff178.lnmtc.cn", "/dzhimg")
		}

		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}
		status := 1
		if v.Status == 2 {
			status = 0
		}
		dzhSchool = append(dzhSchool, &do.AddonsCustomerProSchool{
			Id:         v.Id,
			Logo:       logo,
			Name:       v.Name,
			Address:    v.Address,
			Desc:       v.Desc,
			Content:    v.Content,
			Status:     status,
			CreateTime: createTime,
			UpdateTime: updateTime,
		})
	}

	//插入学校
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProSchool.Table(), dzhSchool, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	var (
		oiMajors  []*do.OiMajors
		dzhMajors []*do.AddonsCustomerProMajors
	)
	//读取专业
	err = customerDao.OiMajors.Ctx(ctx).Where(whereData).Scan(&oiMajors)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	for _, v := range oiMajors {

		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}
		status := 1
		if v.Status == 2 {
			status = 0
		}
		dzhMajors = append(dzhMajors, &do.AddonsCustomerProMajors{
			Id:                v.Id,
			SchoolId:          v.SchoolId,
			Name:              v.Name,
			Amount:            v.Amount,
			PlannedNumbers:    v.PlannedNumbers,
			RegisteredNumbers: v.RegisteredNumbers,
			Status:            status,
			CreateTime:        createTime,
			UpdateTime:        updateTime,
		})
	}

	//插入专业
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProMajors.Table(), dzhMajors, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	var (
		oiTypes  []*do.OiTypes
		dzhTypes []*do.AddonsCustomerProReadtypes
	)
	//读取报读类型
	err = customerDao.OiTypes.Ctx(ctx).Where(whereData).Scan(&oiTypes)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	for _, v := range oiTypes {

		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}
		status := 1
		if v.Status == 2 {
			status = 0
		}
		dzhTypes = append(dzhTypes, &do.AddonsCustomerProReadtypes{
			Id:         v.Id,
			Name:       v.Name,
			Status:     status,
			CreateTime: createTime,
			UpdateTime: updateTime,
		})
	}

	//插入报读类型
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProReadtypes.Table(), dzhTypes, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	var (
		oiDegree  []*do.OiDegree
		dzhDegree []*do.AddonsCustomerProReaddegree
	)
	//读取报读层次
	err = customerDao.OiDegree.Ctx(ctx).Where(whereData).Scan(&oiDegree)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	for _, v := range oiDegree {

		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}
		status := 1
		if v.Status == 2 {
			status = 0
		}
		dzhDegree = append(dzhDegree, &do.AddonsCustomerProReaddegree{
			Id:         v.Id,
			Name:       v.Name,
			Status:     status,
			CreateTime: createTime,
			UpdateTime: updateTime,
		})
	}

	//插入报读层次
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProReaddegree.Table(), dzhDegree, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil

}

// 微信用户
func wxUser(tx gdb.TX) error {

	var (
		oiWechatMember  []*do.OiWechatMember
		dzhWechatMember []*do.AddonsCustomerProWxUser
	)
	whereData := g.Map{
		"deletetime IS NULL": "",
		"id":                 83,
	}

	one, err := customerDao.OiWechatMember.Ctx(ctx).Where(whereData).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	g.Dump(one)

	//读取微信用户
	err = customerDao.OiWechatMember.Ctx(ctx).Where(whereData).Scan(&oiWechatMember)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	for _, v := range oiWechatMember {

		createTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			createTime = gtime.New(gconv.Int(v.Createtime))
		}
		updateTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			updateTime = gtime.New(gconv.Int(v.Updatetime))
		}
		subscribeTime := gtime.Now()
		if gconv.Int(v.Createtime) != 0 {
			subscribeTime = gtime.New(gconv.Int(v.SubscribeTime))
		}
		notify := 1
		if v.Notify == 2 {
			notify = 0
		}
		dzhWechatMember = append(dzhWechatMember, &do.AddonsCustomerProWxUser{
			Id:             v.Id,
			Type:           v.Type,
			UserId:         v.UserId,
			Unionid:        v.Unionid,
			Notify:         notify,
			Openid:         v.Openid,
			Nickname:       v.Nickname,
			Sex:            v.Sex,
			Language:       v.Language,
			Country:        v.Country,
			Province:       v.Province,
			City:           v.City,
			Headimgurl:     v.Headimgurl,
			Subscribe:      v.Subscribe,
			SubscribeTime:  subscribeTime,
			Remark:         v.Remark,
			Groupid:        v.Groupid,
			TagidList:      v.TagidList,
			Privilege:      v.Privilege,
			SubscribeScene: v.SubscribeScene,
			QrScene:        v.QrScene,
			QrSceneStr:     v.QrSceneStr,
			CreateTime:     createTime,
			UpdateTime:     updateTime,
		})
	}

	//插入微信用户
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProWxUser.Table(), dzhWechatMember, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}

// 角色表和用户表迁移
func roleAndUserMove(tx gdb.TX) (upMap g.Map, upgMap g.Map, err error) {

	var (
		groupSlice          []*customerEntity.OiUserGroup
		userSlice           []*customerEntity.OiUser
		baseRoleSlice       []*g.Map
		baseUserSlice       []*g.Map
		baseUserRoleSlice   []*g.Map
		addonsKfSlice       []*g.Map
		groupIdGarray       = garray.NewIntArrayFrom(g.SliceInt{2, 3, 4})
		userProjectMap      = g.Map{}
		userProjectGroupMap = g.Map{}
	)

	//读取角色表
	whereData := g.Map{
		"status":             1,
		"deletetime IS NULL": "",
	}
	err = customerDao.OiUserGroup.Ctx(ctx).Where(whereData).Scan(&groupSlice)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}

	for _, v := range groupSlice {
		if v.Id == 1 {
			continue
		}
		baseRoleSlice = append(baseRoleSlice, &g.Map{
			"id":         gconv.String(v.Id),
			"name":       v.Name,
			"userId":     "1",
			"label":      v.Name,
			"createTime": gtime.New(v.Createtime),
			"updateTime": gtime.New(v.Updatetime),
		})
	}
	//录入角色表
	err = insertDataByQueue(tx, baseDao.BaseSysRole.Table(), baseRoleSlice, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	//////////

	//读取会员表
	err = customerDao.OiUser.Ctx(ctx).Where(whereData).Scan(&userSlice)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}

	for _, v := range userSlice {
		headImg := "/dzhimg/assets/adminking.png"

		if v.Avatar != "" {
			headImg = gstr.Replace(v.Avatar, "http://ff178.lnmtc.cn", "/dzhimg")
		}

		createTime := gtime.Now()
		if v.Createtime != 0 {
			createTime = gtime.New(v.Createtime)
		}

		//会员表
		baseUserSlice = append(baseUserSlice, &g.Map{
			"id":           gconv.String(v.Id),
			"name":         v.Username,
			"nickName":     v.Nickname,
			"username":     v.Mobile,
			"password":     gmd5.MustEncryptString("123456"),
			"phone":        v.Mobile,
			"headImg":      headImg,
			"remark":       v.Desc,
			"departmentId": 1,
			"createTime":   createTime,
			"updateTime":   gtime.New(v.Updatetime),
		})

		//会员和关系表
		baseUserRoleSlice = append(baseUserRoleSlice, &g.Map{
			"id":         dzhcore.NodeSnowflake.Generate().Int64(),
			"userId":     v.Id,
			"roleId":     v.GroupId,
			"createTime": createTime,
			"updateTime": gtime.New(v.Updatetime),
		})

		//客服组表
		if groupIdGarray.Contains(gconv.Int(v.GroupId)) {

			role := 1
			switch v.GroupId {
			case 4:
				role = 1 //成员
			case 3:
				role = 2 //组主管
			case 2:
				role = 3 //项目主管
			}
			status := 1
			if v.AcceptClues == 2 {
				status = 0
			}
			addonsKfSlice = append(addonsKfSlice, &g.Map{
				"id":         dzhcore.NodeSnowflake.Generate().Int64(),
				"name":       v.Username,
				"userId":     v.Id,
				"groupId":    v.ServicesGroupId,
				"projectId":  v.ProjectId,
				"role":       role,
				"status":     status,
				"createTime": createTime,
				"updateTime": gtime.New(v.Updatetime),
			})

			//	项目主管map
			if role == 3 {
				userProjectMap[gconv.String(v.ProjectId)] = gconv.String(v.Id)
			}

			// 客服组主管map
			if role == 2 {
				userProjectGroupMap[gconv.String(v.ServicesGroupId)] = gconv.String(v.Id)
			}

			//会员名称
			userMap[gconv.String(v.Id)] = v.Username
		}

	}

	//录入会员表
	//_, err = tx.Model(baseDao.BaseSysUser.Table()).Data(baseUserSlice).Insert()
	//if err != nil {
	//	g.Log().Error(ctx, err.Error())
	//	return
	//}
	err = insertDataByQueue(tx, baseDao.BaseSysUser.Table(), baseUserSlice, true)
	if err != nil {

		g.Log().Error(ctx, err)
		return
	}

	//录入会员和关系表
	//_, err = tx.Model(baseDao.BaseSysUserRole.Table()).Data(baseUserRoleSlice).Insert()
	//if err != nil {
	//	g.Log().Error(ctx, err.Error())
	//	return
	//}
	err = insertDataByQueue(tx, baseDao.BaseSysUserRole.Table(), baseUserRoleSlice, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	//录入客服组表
	//_, err = tx.Model(customerDao.AddonsCustomerProKf.Table()).Data(addonsKfSlice).Insert()
	//if err != nil {
	//	g.Log().Error(ctx, err.Error())
	//	return
	//}
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProKf.Table(), addonsKfSlice, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return userProjectMap, userProjectGroupMap, nil
}

// 项目表和组迁移
func projectAndGroup(tx gdb.TX, userProjectMap g.Map, userProjectGroupMap g.Map) error {

	var (
		oiProject       []*customerEntity.OiProject
		projectSlice    []*g.Map
		oiServicesGroup []*customerEntity.OiServicesGroup
		projectGroup    []*g.Map
	)

	//读取项目组
	whereData := g.Map{
		"status":             1,
		"deletetime IS NULL": "",
	}
	err := customerDao.OiProject.Ctx(ctx).Where(whereData).Scan(&oiProject)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return err
	}
	for _, v := range oiProject {
		status := 1
		if v.Status == 2 {
			status = 0
		}
		createTime := gtime.Now()
		if v.Createtime != 0 {
			createTime = gtime.New(v.Createtime)
		}
		project := &g.Map{
			"id":            gconv.String(v.Id),
			"name":          v.Name,
			"remark":        v.Desc,
			"status":        status,
			"projectUserId": userProjectMap[gconv.String(v.Id)],
			"createTime":    createTime,
			"updateTime":    gtime.New(v.Updatetime),
		}
		projectSlice = append(projectSlice, project)
	}

	//插入项目组
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProProject.Table(), projectSlice, false)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	//读取客服组
	err = customerDao.OiServicesGroup.Ctx(ctx).Where(whereData).Scan(&oiServicesGroup)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return err
	}

	for _, v := range oiServicesGroup {
		status := 1
		if v.Status == 2 {
			status = 0
		}
		createTime := gtime.Now()
		if v.Createtime != 0 {
			createTime = gtime.New(v.Createtime)
		}
		group := &g.Map{
			"id":          gconv.String(v.Id),
			"projectId":   gconv.String(v.ProjectId),
			"name":        v.Name,
			"status":      status,
			"groupUserId": userProjectGroupMap[gconv.String(v.Id)],
			"createTime":  createTime,
			"updateTime":  gtime.New(v.Updatetime),
		}
		projectGroup = append(projectGroup, group)
	}
	//插入客服组
	err = insertDataByQueue(tx, customerDao.AddonsCustomerProProjectGroup.Table(), projectGroup, false)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil

}

// 数据写入
func insertDataByQueue(tx gdb.TX, table string, insertData interface{}, isUnscoped bool) error {

	m := tx.Model(table).Data(insertData)
	if isUnscoped {
		m = m.Unscoped()
	}
	_, err := m.Batch(1000).OmitEmpty().Insert()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	g.Log().Warningf(ctx, "写入数据：%v", table)
	return nil
}

// 统计
func statistics() {

	schoolPaymentSum, err := customerDao.AddonsCustomerProOrder.Ctx(ctx).Sum("school_payment")
	if err != nil {
		g.Log().Error(ctx, err)
	}

	g.Dump(schoolPaymentSum)

	teamsPaymentSum, err := customerDao.AddonsCustomerProOrder.Ctx(ctx).Sum("teams_payment")
	if err != nil {
		g.Log().Error(ctx, err)
	}
	g.Dump(teamsPaymentSum)

	sum := decimal.NewFromFloat(schoolPaymentSum).Add(decimal.NewFromFloat(teamsPaymentSum))
	g.DumpWithType(sum)

}

func TestClear(t *testing.T) {

	tx := g.DB()
	tables := "addons_customer_pro_clues,addons_customer_pro_clues_record,addons_customer_pro_clues_track,addons_customer_pro_kf,addons_customer_pro_majors,addons_customer_pro_order,addons_customer_pro_project,addons_customer_pro_project_group,addons_customer_pro_readdegree,addons_customer_pro_readtypes,addons_customer_pro_school,addons_customer_pro_wx_user,base_sys_user,base_sys_role,base_sys_user_role,addons_customer_pro_chatContent"
	tableSlice := gstr.Split(tables, ",")
	for _, table := range tableSlice {
		if table == "base_sys_role" || table == "base_sys_user" || table == "base_sys_user_role" {
			if table == "base_sys_role" {
				_, err := tx.Model(table).WhereNot("id", "1").Unscoped().Delete()
				if err != nil {
					g.Log().Error(ctx, err)
					return
				}
			}
			if table == "base_sys_user" {
				_, err := tx.Model(table).Where("id NOT IN (?)", g.SliceStr{"1152921504606846975", "1152921504606846976"}).Unscoped().Delete()
				if err != nil {
					g.Log().Error(ctx, err)
					return
				}
			}
			if table == "base_sys_user_role" {
				_, err := tx.Model(table).Where("id NOT IN (?)", g.SliceStr{"1152921504606846977", "1152921504606846978"}).Unscoped().Delete()
				if err != nil {
					g.Log().Error(ctx, err)
					return
				}
			}

		} else {
			_, err := tx.Model(table).Where("id IS NOT NULL").Unscoped().Delete()
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
		}
		g.Log().Warningf(ctx, "清空：%v", table)
	}
}

func TestJoinDaoSql(t *testing.T) {

	err := dzhcore.CacheManager.Set(ctx, "admin:test:1", "1111111", 0)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	//设置查询结果放到redis
	redisCache := gcache.NewAdapterRedis(g.Redis("core"))
	customerDao.AddonsCustomerProClues.DB().GetCache().SetAdapter(redisCache)

	m := customerDao.AddonsCustomerProClues.Ctx(ctx)
	m = m.OrderDesc("createTime")

	for i := 0; i < 2; i++ {
		result, err := m.Cache(gdb.CacheOption{
			Duration: time.Hour,
			Name:     "clues",
			Force:    false,
		}).Limit(1).All()
		if err != nil {
			g.Log().Error(ctx, err)
			panic(err)
		}
		g.Log().Infof(ctx, gjson.MustEncodeString(result))
	}

	time.Sleep(time.Second * 5)

	//清理指定key的缓存
	//err = customerDao.AddonsCustomerProClues.DB().GetCore().ClearCache(ctx, "clues")
	_, err = dzhcore.CacheManager.Remove(ctx, "SelectCache:clues")
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}

	for i := 0; i < 2; i++ {
		result, err := m.Cache(gdb.CacheOption{
			Duration: time.Hour,
			Name:     "vip",
			Force:    false,
		}).Limit(1).All()
		if err != nil {
			g.Log().Error(ctx, err)
			panic(err)
		}
		g.Log().Infof(ctx, gjson.MustEncodeString(result))
	}

}

func TestDaoSql(t *testing.T) {

	var (
		db  = g.DB()
		ctx = gctx.New()
	)
	// 开启调试模式，以便于记录所有执行的SQL
	db.SetDebug(true)
	id := "1825099577766711296"
	//id := dzhcore.NodeSnowflake.Generate().String()
	// 写入测试数据
	//_, err := g.Model("addons_member_manage").Ctx(ctx).Data(g.Map{
	//	"id":       id,
	//	"username": "john",
	//	"password": 123456,
	//}).Insert()
	// 执行2次查询并将查询结果缓存1小时，并可执行缓存名称(可选)
	for i := 0; i < 2; i++ {
		r, _ := g.Model("addons_member_manage").Ctx(ctx).Cache(gdb.CacheOption{
			Duration: time.Hour,
			Name:     "vip-user",
			Force:    false,
		}).All()
		g.Log().Debug(ctx, r.MapKeyInt("id"))
	}
	// 执行更新操作，并清理指定名称的查询缓存
	_, err := g.Model("addons_member_manage").Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "vip-user",
		Force:    false,
	}).Data(gdb.Map{"username": "smith"}).Where("id", id).Update()
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 再次执行查询，启用查询缓存特性
	r, _ := g.Model("addons_member_manage").Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     "vip-user",
		Force:    false,
	}).All()

	g.Log().Debug(ctx, r.MapKeyInt("id"))

}

func TestReg(t *testing.T) {

	openReg := `^(/admin/.*/open/.*|/app/.*/open/.*)$`

	testCases := []string{
		"/admin/something/open/page",
		"/admin/123/open/list",
		"/app/something/open/page",
		"/app/anything/open/list",
		"/admin/somethingelse/page",
		"/app/somethingelse/list",
	}

	for _, testCase := range testCases {
		g.Log().Info(ctx, gregex.IsMatch(openReg, []byte(testCase)))
	}

}

func TestAdmin(t *testing.T) {

	type Admin struct {
		IsRefresh       bool    `json:"isRefresh"`
		RoleIds         []int64 `json:"roleIds"`
		Username        string  `json:"username"`
		UserId          int64   `json:"userId"`
		PasswordVersion *int32  `json:"passwordVersion"`
	}
	admin := &Admin{}
	jsonStr := `{"isRefresh":false,"roleIds":[1],"username":"admin","userId":1152921504606846975,"passwordVersion":3,"exp":1724522584,"iat":3917784}`

	err := gjson.New(jsonStr).Scan(admin)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	g.Dump(admin)

	admin2 := &Admin{}
	err = json.Unmarshal([]byte(jsonStr), admin2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	g.Dump(admin2)
}
