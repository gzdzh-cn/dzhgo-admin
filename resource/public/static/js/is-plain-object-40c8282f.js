import{i as t}from"./isobject-c886ab54.js";
/*!
 * is-plain-object <https://github.com/jonschlinkert/is-plain-object>
 *
 * Copyright (c) 2014-2017, Jon Schlinkert.
 * Released under the MIT License.
 */var o=t;function r(t){return!0===o(t)&&"[object Object]"===Object.prototype.toString.call(t)}var e=function(t){var o,e;return!1!==r(t)&&("function"==typeof(o=t.constructor)&&(!1!==r(e=o.prototype)&&!1!==e.hasOwnProperty("isPrototypeOf")))};export{e as i};
