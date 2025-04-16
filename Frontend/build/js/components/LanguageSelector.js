"use strict";

function _slicedToArray(r, e) { return _arrayWithHoles(r) || _iterableToArrayLimit(r, e) || _unsupportedIterableToArray(r, e) || _nonIterableRest(); }
function _nonIterableRest() { throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); }
function _unsupportedIterableToArray(r, a) { if (r) { if ("string" == typeof r) return _arrayLikeToArray(r, a); var t = {}.toString.call(r).slice(8, -1); return "Object" === t && r.constructor && (t = r.constructor.name), "Map" === t || "Set" === t ? Array.from(r) : "Arguments" === t || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t) ? _arrayLikeToArray(r, a) : void 0; } }
function _arrayLikeToArray(r, a) { (null == a || a > r.length) && (a = r.length); for (var e = 0, n = Array(a); e < a; e++) n[e] = r[e]; return n; }
function _iterableToArrayLimit(r, l) { var t = null == r ? null : "undefined" != typeof Symbol && r[Symbol.iterator] || r["@@iterator"]; if (null != t) { var e, n, i, u, a = [], f = !0, o = !1; try { if (i = (t = t.call(r)).next, 0 === l) { if (Object(t) !== t) return; f = !1; } else for (; !(f = (e = i.call(t)).done) && (a.push(e.value), a.length !== l); f = !0); } catch (r) { o = !0, n = r; } finally { try { if (!f && null != t["return"] && (u = t["return"](), Object(u) !== u)) return; } finally { if (o) throw n; } } return a; } }
function _arrayWithHoles(r) { if (Array.isArray(r)) return r; }
function LanguageSelector() {
  try {
    var _React$useState = React.useState(false),
      _React$useState2 = _slicedToArray(_React$useState, 2),
      isOpen = _React$useState2[0],
      setIsOpen = _React$useState2[1];
    var _React$useContext = React.useContext(LanguageContext),
      language = _React$useContext.language,
      setLanguage = _React$useContext.setLanguage;
    var languages = {
      en: 'English',
      zh: '中文'
    };
    return /*#__PURE__*/React.createElement("div", {
      className: "language-selector",
      "data-name": "language-selector"
    }, /*#__PURE__*/React.createElement("div", {
      className: "flex items-center cursor-pointer",
      onClick: function onClick() {
        return setIsOpen(!isOpen);
      },
      "data-name": "language-selector-trigger"
    }, /*#__PURE__*/React.createElement("i", {
      className: "fas fa-globe mr-2"
    }), /*#__PURE__*/React.createElement("span", null, languages[language])), isOpen && /*#__PURE__*/React.createElement("div", {
      className: "language-menu",
      "data-name": "language-menu"
    }, Object.entries(languages).map(function (_ref) {
      var _ref2 = _slicedToArray(_ref, 2),
        code = _ref2[0],
        name = _ref2[1];
      return /*#__PURE__*/React.createElement("div", {
        key: code,
        className: "px-4 py-2 hover:bg-gray-100 cursor-pointer ".concat(language === code ? 'text-blue-600' : 'text-gray-700'),
        onClick: function onClick() {
          setLanguage(code);
          setIsOpen(false);
        },
        "data-name": "language-option-".concat(code)
      }, name);
    })));
  } catch (error) {
    console.error('LanguageSelector component error:', error);
    reportError(error);
    return null;
  }
}