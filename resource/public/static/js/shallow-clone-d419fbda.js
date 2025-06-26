import{k as r}from"./kind-of-d78bf752.js";
/*!
 * shallow-clone <https://github.com/jonschlinkert/shallow-clone>
 *
 * Copyright (c) 2015-present, Jon Schlinkert.
 * Released under the MIT License.
 */const e=Symbol.prototype.valueOf,t=r;var a=function(r,a){switch(t(r)){case"array":return r.slice();case"object":return Object.assign({},r);case"date":return new r.constructor(Number(r));case"map":return new Map(r);case"set":return new Set(r);case"buffer":return function(r){const e=r.length,t=Buffer.allocUnsafe?Buffer.allocUnsafe(e):Buffer.from(e);return r.copy(t),t}(r);case"symbol":return function(r){return e?Object(e.call(r)):{}}(r);case"arraybuffer":return function(r){const e=new r.constructor(r.byteLength);return new Uint8Array(e).set(new Uint8Array(r)),e}(r);case"float32array":case"float64array":case"int16array":case"int32array":case"int8array":case"uint16array":case"uint32array":case"uint8clampedarray":case"uint8array":return function(r,e){return new r.constructor(r.buffer,r.byteOffset,r.length)}(r);case"regexp":return function(r){const e=void 0!==r.flags?r.flags:/\w+$/.exec(r)||void 0,t=new r.constructor(r.source,e);return t.lastIndex=r.lastIndex,t}(r);case"error":return Object.create(r);default:return r}};export{a as s};
