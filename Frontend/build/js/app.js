"use strict";

var _react = _interopRequireDefault(require("react"));
var _client = _interopRequireDefault(require("react-dom/client"));
var _LanguageContext = require("./path/to/LanguageContext");
var _Home = _interopRequireDefault(require("./path/to/Home"));
var _AboutUs = _interopRequireDefault(require("./path/to/AboutUs"));
var _Login = _interopRequireDefault(require("./path/to/Login"));
var _Register = _interopRequireDefault(require("./path/to/Register"));
var _ForgotPassword = _interopRequireDefault(require("./path/to/ForgotPassword"));
var _ResetPassword = _interopRequireDefault(require("./path/to/ResetPassword"));
function _interopRequireDefault(e) { return e && e.__esModule ? e : { "default": e }; }
function _slicedToArray(r, e) { return _arrayWithHoles(r) || _iterableToArrayLimit(r, e) || _unsupportedIterableToArray(r, e) || _nonIterableRest(); }
function _nonIterableRest() { throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); }
function _unsupportedIterableToArray(r, a) { if (r) { if ("string" == typeof r) return _arrayLikeToArray(r, a); var t = {}.toString.call(r).slice(8, -1); return "Object" === t && r.constructor && (t = r.constructor.name), "Map" === t || "Set" === t ? Array.from(r) : "Arguments" === t || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t) ? _arrayLikeToArray(r, a) : void 0; } }
function _arrayLikeToArray(r, a) { (null == a || a > r.length) && (a = r.length); for (var e = 0, n = Array(a); e < a; e++) n[e] = r[e]; return n; }
function _iterableToArrayLimit(r, l) { var t = null == r ? null : "undefined" != typeof Symbol && r[Symbol.iterator] || r["@@iterator"]; if (null != t) { var e, n, i, u, a = [], f = !0, o = !1; try { if (i = (t = t.call(r)).next, 0 === l) { if (Object(t) !== t) return; f = !1; } else for (; !(f = (e = i.call(t)).done) && (a.push(e.value), a.length !== l); f = !0); } catch (r) { o = !0, n = r; } finally { try { if (!f && null != t["return"] && (u = t["return"](), Object(u) !== u)) return; } finally { if (o) throw n; } } return a; } }
function _arrayWithHoles(r) { if (Array.isArray(r)) return r; }
function App() {
  try {
    var _React$useState = _react["default"].useState('home'),
      _React$useState2 = _slicedToArray(_React$useState, 2),
      currentPage = _React$useState2[0],
      setCurrentPage = _React$useState2[1];
    var _React$useState3 = _react["default"].useState('en'),
      _React$useState4 = _slicedToArray(_React$useState3, 2),
      language = _React$useState4[0],
      setLanguage = _React$useState4[1];
    var handleNavigate = function handleNavigate(page) {
      setCurrentPage(page);
    };
    var renderPage = function renderPage() {
      switch (currentPage) {
        case 'home':
          return /*#__PURE__*/_react["default"].createElement(_Home["default"], {
            onNavigate: handleNavigate
          });
        case 'about':
          return /*#__PURE__*/_react["default"].createElement(_AboutUs["default"], {
            onNavigate: handleNavigate
          });
        case 'login':
          return /*#__PURE__*/_react["default"].createElement(_Login["default"], {
            onNavigate: handleNavigate
          });
        case 'register':
          return /*#__PURE__*/_react["default"].createElement(_Register["default"], {
            onNavigate: handleNavigate
          });
        case 'forgot-password':
          return /*#__PURE__*/_react["default"].createElement(_ForgotPassword["default"], {
            onNavigate: handleNavigate
          });
        case 'reset-password':
          return /*#__PURE__*/_react["default"].createElement(_ResetPassword["default"], {
            onNavigate: handleNavigate
          });
        default:
          return /*#__PURE__*/_react["default"].createElement(_Home["default"], {
            onNavigate: handleNavigate
          });
      }
    };
    return /*#__PURE__*/_react["default"].createElement(_LanguageContext.LanguageContext.Provider, {
      value: {
        language: language,
        setLanguage: setLanguage
      }
    }, /*#__PURE__*/_react["default"].createElement("div", {
      className: "min-h-screen bg-gray-100",
      "data-name": "app-container"
    }, renderPage()));
  } catch (error) {
    console.error('App component error:', error);
    reportError(error);
    return null;
  }
}
var root = _client["default"].createRoot(document.getElementById('root'));
root.render(/*#__PURE__*/_react["default"].createElement(App, null));