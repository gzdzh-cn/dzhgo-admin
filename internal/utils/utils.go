package utils

import (
	"context"
	"dzhgo/internal/defineType"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh-cn/dzhcore"
	"github.com/shopspring/decimal"
	"gopkg.in/gomail.v2"
	"os"
	"runtime"
	"strings"
	"time"
)

// 发送邮件
func SentEmail(ctx context.Context, content string, subject string, addressHeader string, config g.Map) {

	g.Log().Debug(ctx, "util SentEmail")
	os.Setenv("GODEBUG", "tlsrsakex=1")

	// 设置 SMTP 服务器的认证信息
	smtp := gconv.String(config["smtp"])
	smtpPort := 465
	senderEmail := gconv.String(config["smtpEmail"])
	senderPassword := gconv.String(config["smtpPass"])

	body := content
	// 邮件内容
	toEmail := gconv.String(config["remindEmail"])
	toEmails := strings.Split(toEmail, "|")

	m := gomail.NewMessage()
	m.SetHeader("To", toEmails...)
	m.SetHeader("Subject", subject)
	m.SetAddressHeader("From", senderEmail, addressHeader)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtp, smtpPort, senderEmail, senderPassword)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return
	}
	g.Log().Info(ctx, "邮件发送成功")

}

// 多行文本转一行
func StrTranLine(jsonData string) string {
	// 移除换行符和制表符
	oneLine := strings.ReplaceAll(jsonData, "\n", "")
	oneLine = strings.ReplaceAll(oneLine, "\t", "")
	oneLine = strings.ReplaceAll(oneLine, "  ", "") // 可根据需要移除多余空格
	return oneLine
}

// 日式打印运行时间
func StdOutLog(ctx context.Context, startTime time.Time, memStatsStart runtime.MemStats) {

	var (
		r           = g.RequestFromCtx(ctx)
		ctxId       = gctx.CtxId(r.GetCtx()) //获取当前请求的ctxid
		elapsedTime = time.Since(startTime)  // 请求处理时间
		outLogger   defineType.OutputsForLogger
		path        string
		stdout      bool
		memStatsEnd runtime.MemStats // 记录结束内存状态

	)

	// 根据处理时间计算吞吐率
	throughput := 1.0 / elapsedTime.Seconds() //（秒）
	runtime.ReadMemStats(&memStatsEnd)
	// 计算内存消耗
	memUsed := memStatsEnd.Alloc - memStatsStart.Alloc

	file := g.Cfg().MustGet(ctx, "modules.base.middleware.runLogger.file").String()
	path = dzhcore.GetCfgWithDefault(ctx, "modules.base.middleware.runLogger.path", g.NewVar(false)).String()
	stdout = dzhcore.GetCfgWithDefault(ctx, "modules.base.middleware.runLogger.stdout", g.NewVar(false)).Bool()
	outLogger = defineType.OutputsForLogger{
		Time:       time.Now(),
		Host:       r.Host,
		RequestURI: r.RequestURI,
		Params:     gjson.MustEncodeString(r.GetMap()),
		RunTime:    float64(elapsedTime.Nanoseconds()) / 1e9,
		File:       file,
		FileRule:   gstr.Trim(gstr.Split(file, ".")[0], "{}"),
		RotateSize: dzhcore.GetCfgWithDefault(ctx, "modules.base.middleware.runLogger.rotateSize", g.NewVar(100000)).String(),
		Stdout:     stdout,
		Path:       path,
		Throughput: throughput,
		MemUsed:    memUsed,
	}

	fileName := fmt.Sprintf("%s.log", gtime.Now().Format(outLogger.FileRule))
	tempFile := fmt.Sprintf("%v%v", path, fileName)

	throughputStringFixed := decimal.NewFromFloat(outLogger.Throughput).StringFixed(2)

	logSlice := g.SliceStr{
		fmt.Sprintf("[ %s ] %s OPTIONS %s\n", outLogger.Time, outLogger.Host, outLogger.RequestURI),
		fmt.Sprintf("[ 运行时间：%vs ] [TraceId：%v ] [ 吞吐率：%vreq/s ] [ 内存消耗：%v ]\n", outLogger.RunTime, ctxId, throughputStringFixed, humanize.Bytes(outLogger.MemUsed)),
		fmt.Sprintf("[ info ] [ PARAM ] [ %v ]\n", StrTranLine(outLogger.Params)),
	}

	//超过容量就切割
	byteSize, _ := humanize.ParseBytes(outLogger.RotateSize)
	if gfile.Size(tempFile) > int64(byteSize) {
		endTime := gtime.Now().Format("H-i-s")
		dstPath := tempFile + "." + endTime
		gfile.Rename(tempFile, dstPath)
	}

	//写入到日志
	for _, log := range logSlice {
		gfile.PutContentsAppend(tempFile, log)
	}
	gfile.PutContentsAppend(tempFile, "----------------------------------------\n")

	//打印到控制台
	if stdout {
		for _, log := range logSlice {
			g.Log().Info(ctx, StrTranLine(log))
		}
		g.Log().Info(ctx, "----------------------------------------")
	}
}
