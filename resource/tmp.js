(function (t) {
    function e(e) {
        for (var i, s, r = e[0], c = e[1], l = e[2], d = 0, p = []; d < r.length; d++) s = r[d], a[s] && p.push(a[s][0]), a[s] = 0;
        for (i in c) Object.prototype.hasOwnProperty.call(c, i) && (t[i] = c[i]);
        u && u(e);
        while (p.length) p.shift()();
        return o.push.apply(o, l || []), n()
    }

    function n() {
        for (var t, e = 0; e < o.length; e++) {
            for (var n = o[e], i = !0, r = 1; r < n.length; r++) {
                var c = n[r];
                0 !== a[c] && (i = !1)
            }
            i && (o.splice(e--, 1), t = s(s.s = n[0]))
        }
        return t
    }

    var i = {}, a = {index: 0}, o = [];

    function s(e) {
        if (i[e]) return i[e].exports;
        var n = i[e] = {i: e, l: !1, exports: {}};
        return t[e].call(n.exports, n, n.exports, s), n.l = !0, n.exports
    }

    s.m = t, s.c = i, s.d = function (t, e, n) {
        s.o(t, e) || Object.defineProperty(t, e, {enumerable: !0, get: n})
    }, s.r = function (t) {
        "undefined" !== typeof Symbol && Symbol.toStringTag && Object.defineProperty(t, Symbol.toStringTag, {value: "Module"}), Object.defineProperty(t, "__esModule", {value: !0})
    }, s.t = function (t, e) {
        if (1 & e && (t = s(t)), 8 & e) return t;
        if (4 & e && "object" === typeof t && t && t.__esModule) return t;
        var n = Object.create(null);
        if (s.r(n), Object.defineProperty(n, "default", {
            enumerable: !0,
            value: t
        }), 2 & e && "string" != typeof t) for (var i in t) s.d(n, i, function (e) {
            return t[e]
        }.bind(null, i));
        return n
    }, s.n = function (t) {
        var e = t && t.__esModule ? function () {
            return t["default"]
        } : function () {
            return t
        };
        return s.d(e, "a", e), e
    }, s.o = function (t, e) {
        return Object.prototype.hasOwnProperty.call(t, e)
    }, s.p = "";
    var r = window["webpackJsonp"] = window["webpackJsonp"] || [], c = r.push.bind(r);
    r.push = e, r = r.slice();
    for (var l = 0; l < r.length; l++) e(r[l]);
    var u = c;
    o.push([0, "chunk-vendors"]), n()
})({
    0: function (t, e, n) {
        t.exports = n("56d7")
    }, "0815": function (t, e, n) {
        "use strict";
        var i = n("ae5a"), a = n.n(i);
        a.a
    }, 2856: function (t, e, n) {
    }, "56d7": function (t, e, n) {
        "use strict";
        n.r(e);
        n("cadf"), n("551c"), n("097d");
        var i = n("2b0e"), a = function () {
            var t = this, e = t.$createElement, n = t._self._c || e;
            return n("div", {attrs: {id: "app"}}, [n("div", {staticClass: "main-body"}, [n("div", {staticClass: "main-head"}, [t._v("\n      邀请码\n    ")]), n("div", {staticClass: "main-panel"}, [t._l(t.panmel_data, function (e, i) {
                return [n("div", {
                    key: i, staticClass: "panmel-item", on: {
                        click: function (n) {
                            t.clich_select_panel(e)
                        }
                    }
                }, [t._v("\n          " + t._s(e.name) + "\n          "), n("div", {
                    directives: [{
                        name: "show",
                        rawName: "v-show",
                        value: e.isShow,
                        expression: "item.isShow"
                    }]
                })])]
            })], 2), n("div", {staticClass: "main-content"}, [n("div", {
                staticStyle: {
                    position: "absolute",
                    width: "22px",
                    height: "22px",
                    "border-radius": "50%",
                    "background-color": "#ffa60c",
                    top: "-11px",
                    left: "-11px"
                }
            }), n("div", {
                staticStyle: {
                    position: "absolute",
                    width: "22px",
                    height: "22px",
                    "border-radius": "50%",
                    "background-color": "#fdc607",
                    top: "-11px",
                    right: "-11px"
                }
            }), n("div", {
                staticStyle: {
                    position: "absolute",
                    width: "12px",
                    height: "12px",
                    "border-radius": "50%",
                    "background-color": "#ffa60c",
                    top: "46.3%",
                    left: "-6px"
                }
            }), n("div", {
                staticStyle: {
                    position: "absolute",
                    width: "12px",
                    height: "12px",
                    "border-radius": "50%",
                    "background-color": "#fdc607",
                    top: "46.3%",
                    right: "-6px"
                }
            }), n("div", {
                staticStyle: {
                    position: "absolute",
                    width: "22px",
                    height: "22px",
                    "border-radius": "50%",
                    "background-color": "#ffa60c",
                    bottom: "-11px",
                    left: "-11px"
                }
            }), n("div", {
                staticStyle: {
                    position: "absolute",
                    width: "22px",
                    height: "22px",
                    "border-radius": "50%",
                    "background-color": "#fdc607",
                    bottom: "-11px",
                    right: "-11px"
                }
            }), n("div", {staticClass: "first-item"}, [n("first-pannel", {
                directives: [{
                    name: "show",
                    rawName: "v-show",
                    value: t.pannel_show,
                    expression: "pannel_show"
                }], attrs: {res: t.resultApi}
            }), n("second-pannel", {
                directives: [{
                    name: "show",
                    rawName: "v-show",
                    value: !t.pannel_show,
                    expression: "!pannel_show"
                }], on: {clickGenerateCode: t.request_code}
            })], 1), t._m(0)]), t._m(1)])])
        }, o = [function () {
            var t = this, e = t.$createElement, i = t._self._c || e;
            return i("div", {staticClass: "second-item"}, [i("div", {staticClass: "QR-code"}, [i("img", {
                attrs: {
                    src: n("a2be"),
                    alt: ""
                }
            })]), i("div", [t._v("扫码下载全球付")]), i("div", [i("span", [t._v("全球付")]), t._v("为您的资产保驾护航")])])
        }, function () {
            var t = this, e = t.$createElement, n = t._self._c || e;
            return n("div", {staticClass: "active-content"}, [n("div", [t._v("活动规则")]), n("div", [t._v("\n          1.每人总共可邀请10位好友加入环球付。"), n("br"), t._v("\n          2.被邀请人在创建钱包并进入钱包后需要在“我-邀请码-填写邀请码”页面输入邀请人的“邀请码”，每位用户可且只能被邀请一次。"), n("br"), t._v("\n          3.成功邀请后，邀请人和被邀请人均可获得2GPC。"), n("br"), n("span", [t._v("温馨提示：圣诞节前后将开放GPC兑换ETH活动，各位请保存好你们的GPC。")])]), n("div", [t._v("活动最终解释权归全球付所有")]), n("div")])
        }], s = (n("ac6a"), n("28a5"), n("386d"), function () {
            var t = this, e = t.$createElement, n = t._self._c || e;
            return n("div", {staticClass: "first-pannel"}, [n("div", [t._v("我的邀请码")]), n("div", [t._v(t._s(t.res.invitCode))]), n("div", [n("button", {
                staticClass: "btn",
                attrs: {"data-clipboard-text": t.copyMsg}
            }, [t._v("复制")])]), n("div", [t._v("剩余邀请次数" + t._s(t.res.times) + "次")]), t._m(0)])
        }), r = [function () {
            var t = this, e = t.$createElement, n = t._self._c || e;
            return n("div", [n("div", [t._v("\n        你已成功领取"), n("span", [t._v("1.8GPC"), n("br"), t._v("1GPC≈0.005ETH")])])])
        }], c = n("b311"), l = n.n(c), u = {
            props: ["res"], data: function () {
                return {remateNum: 10, copyMsg: "", clipboard: void 0}
            }, mounted: function () {
                this.clipboard = new l.a(".btn"), this.copyMsg = "即刻加入环球付，赢取GPC（在圣诞之夜领取ETH神秘大奖），邀请码:" + this.res.invitCode + "，次数有限，赶快行动吧!～https://www.g-pay.io"
            }
        }, d = u, p = (n("f536"), n("2877")), A = Object(p["a"])(d, s, r, !1, null, "14e71ee2", null);
        A.options.__file = "firstPannel.vue";
        var v = A.exports, f = function () {
            var t = this, e = t.$createElement, n = t._self._c || e;
            return n("div", {staticClass: "second-pannel"}, [n("div", [t._v("填写邀请码")]), n("div", [n("input", {
                directives: [{
                    name: "model",
                    rawName: "v-model",
                    value: t.inputVal,
                    expression: "inputVal"
                }], domProps: {value: t.inputVal}, on: {
                    input: function (e) {
                        e.target.composing || (t.inputVal = e.target.value)
                    }
                }
            })]), n("div", {on: {click: t.click_determine}}, [t._v("确定")]), t._m(0)])
        }, g = [function () {
            var t = this, e = t.$createElement, n = t._self._c || e;
            return n("div", [n("div", [t._v("\n        注意：每个用户且只能填写"), n("span", [t._v("一次")]), t._v("邀请码"), n("span", [t._v("即可领取GPC")])])])
        }], h = {
            data: function () {
                return {inputVal: ""}
            }, methods: {
                click_determine: function () {
                    this.$emit("clickGenerateCode", {inputCode: this.inputVal})
                }
            }
        }, m = h, w = (n("0815"), Object(p["a"])(m, f, g, !1, null, "441a043b", null));
        w.options.__file = "secondPannel.vue";
        var b = w.exports, _ = {
            name: "app", data: function () {
                return {
                    panmel_data: [{name: "我的邀请码", isShow: !0, id: "first"}, {
                        name: "填写邀请码",
                        isShow: !1,
                        id: "second"
                    }], pannel_show: !0, resultApi: "", deviceType: null
                }
            }, created: function () {
                if ("" != window.location.search) {
                    var t = window.location.search.split("=")[1];
                    this.deviceType = t
                }
            }, methods: {
                clich_select_panel: function (t) {
                    this.panmel_data.forEach(function (t) {
                        t.isShow = !1
                    }), t.isShow = !0, this.pannel_show = !this.pannel_show
                }, myInvitationCode: function (t) {
                    this.resultApi = t || null
                }, request_code: function (t) {
                    OCObject.startTradeInvitationCode(t.inputCode)
                }
            }, components: {firstPannel: v, secondPannel: b}
        }, E = _, B = (n("5c0b"), Object(p["a"])(E, a, o, !1, null, null, null));
        B.options.__file = "App.vue";
        var x = B.exports;
        i["a"].config.productionTip = !1, new i["a"]({
            render: function (t) {
                return t(x)
            }
        }).$mount("#app")
    }, "5c0b": function (t, e, n) {
        "use strict";
        var i = n("2856"), a = n.n(i);
        a.a
    }, a2be: function (t, e) {
        t.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQQAAAEECAIAAABBat1dAAAE/ElEQVR4nO3dwW7jNhRA0aSY///ldHPRnVRIfSEl95xta8mJc0HAb0h9//z8fAFfX3/tfgPwFGKAiAEiBogYIGKAiAHy5+g/fH9/r3wf/92Cgcng72Tw3c5+Ulff2Cf9nVgZIGKAiAEiBogYIGKAiAFyOGc4snf/w+C32jcudfKzX73ayf8/+Bu+camjN3bjUq/7U7EyQMQAEQNEDBAxQMQAEQNEDJDLQ7cTH7/3Zc2k7KrXHQP3zL+TLysD/EMMEDFAxAARA0QMEDFAJucM73Jjp87g5p7BbTezXje1GGRlgIgBIgaIGCBigIgBIgbI588Zjr44n92c8H/+ev5jWBkgYoCIASIGiBggYoCIASIGyOTQ7V2DpzXbawZHfjcudfUWJ1fb+zyhNawMEDFAxAARA0QMEDFAxAC5PGdYc5TVoMETwQYPEbvhme/qxOv+VKwMEDFAxAARA0QMEDFAxAARA+T7sTstNnrXk3t8glOsDBAxQMQAEQNEDBAxQMQAWTFnWHP21oJL7b3LmlPPBu9+w4JxzQkrA0QMEDFAxAARA0QMEDFADucMs4/AuOqZB2YNnju299d7YnDXxIIpx41nvpywMkDEABEDRAwQMUDEABEDRAyQlx0idmPIsma8tWBWtWbwdGTNPiGbe+ARxAARA0QMEDFAxAARA+TP4LUWfEm85iWPPapso73zqDWTHysDRAwQMUDEABEDRAwQMUDEADkcuu19Hv3eR87sfQzP0V0G7/7YPVJXX+JEPfgVYoCIASIGiBggYoCIATJ5iNjGg7ROXnJk8DE8N6w5+WuB2ZHF1ZfMfohWBogYIGKAiAEiBogYIGKAXN7PMGjvv6rf+03/4Bfke88pm/2mf+/PbmWAiAEiBogYIGKAiAEiBogYIJef3LP39Ku9G1/WjLee+YycGx/u4KRszRFmVgaIGCBigIgBIgaIGCBigFw+RGzv9pq9jwt57K6Uq9YcoLbm5K9BVgaIGCBigIgBIgaIGCBigEw+FH3v40L2fnN/ZM1Y5uql9nrsg1qsDBAxQMQAEQNEDBAxQMQAEQPk8iFie2dbg/Oa2aOsBh9tP3WL7W684b1zWysDRAwQMUDEABEDRAwQMUAODxFbsCtlzU6OBadinVztmUOA2Z998C5X7z77J2RlgIgBIgaIGCBigIgBIgaIGCCXN/fs9dhT5RbsSrlxqRsGp4TP3Admcw/8OzFAxAARA0QMEDFAxAA53NzzMdZ82z14hNngHqnBn33BTp0bbO6BXyEGiBggYoCIASIGiBggh/sZnnn61Ym9A5O9T9m4epc1mxae+YZPWBkgYoCIASIGiBggYoCIASIGyOVDxN412/q6Ncf5mA1PC2ZVa44wW/OJWBkgYoCIASIGiBggYoCIATL5sJJnHg71zI0sr3smyDPPYptlZYCIASIGiBggYoCIASIGiBggk0O3dxk87+1r966Uq9a8q9c9asjKABEDRAwQMUDEABEDRAyQz58zDB4iNvuSqwYPUBu8y+wtBvdI3ZhyWBkgYoCIASIGiBggYoCIATI5Z3jmv90/sndzwoIJwMlL1nxSN2Yvg1OOG78uKwNEDBAxQMQAEQNEDBAxQMQAuTx0W/MElwUGt5icXG3BnG7vM3Ue+0Qfh4jBfWKAiAEiBogYIGKAiAHy/a4dOfB7rAwQMUDEABEDRAwQMUDEABED5G8DQCQmfqVkNAAAAABJRU5ErkJggg=="
    }, ae5a: function (t, e, n) {
    }, c897: function (t, e, n) {
    }, f536: function (t, e, n) {
        "use strict";
        var i = n("c897"), a = n.n(i);
        a.a
    }
});
//# sourceMappingURL=index.9eb1800c.js.map