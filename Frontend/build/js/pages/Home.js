"use strict";

function Home(_ref) {
  var onNavigate = _ref.onNavigate;
  try {
    var _React$useContext = React.useContext(LanguageContext),
      language = _React$useContext.language;
    var t = translations[language].home;
    var features = [{
      icon: 'fa-chart-line',
      text: t.features.marketData
    }, {
      icon: 'fa-money-bill-trend-up',
      text: t.features.trading
    }, {
      icon: 'fa-wallet',
      text: t.features.management
    }, {
      icon: 'fa-brain',
      text: t.features.ai
    }];
    return /*#__PURE__*/React.createElement("div", {
      className: "hero-container",
      "data-name": "home-container"
    }, /*#__PURE__*/React.createElement("div", {
      className: "hero-overlay"
    }), /*#__PURE__*/React.createElement("div", {
      className: "hero-content h-full"
    }, /*#__PURE__*/React.createElement("nav", {
      className: "flex justify-end p-6 relative z-10",
      "data-name": "home-nav"
    }, /*#__PURE__*/React.createElement("div", {
      className: "flex items-center space-x-6"
    }, /*#__PURE__*/React.createElement(LanguageSelector, null), /*#__PURE__*/React.createElement("button", {
      onClick: function onClick() {
        return onNavigate('about');
      },
      className: "text-white hover:text-blue-300 transition-colors",
      "data-name": "about-us-button"
    }, t.aboutUs))), /*#__PURE__*/React.createElement("div", {
      className: "flex flex-col items-center justify-center h-[calc(100%-80px)] px-4 text-center",
      "data-name": "hero-content"
    }, /*#__PURE__*/React.createElement("h1", {
      className: "text-4xl md:text-6xl font-bold text-white mb-4",
      "data-name": "hero-title"
    }, t.title), /*#__PURE__*/React.createElement("p", {
      className: "text-xl md:text-2xl text-white mb-8",
      "data-name": "hero-subtitle"
    }, t.subtitle), /*#__PURE__*/React.createElement("div", {
      className: "grid grid-cols-1 md:grid-cols-2 gap-4 mb-12",
      "data-name": "features-grid"
    }, features.map(function (feature, index) {
      return /*#__PURE__*/React.createElement("div", {
        key: index,
        className: "feature-item",
        "data-name": "feature-".concat(index)
      }, /*#__PURE__*/React.createElement("i", {
        className: "fas ".concat(feature.icon, " feature-icon")
      }), /*#__PURE__*/React.createElement("span", null, feature.text));
    })), /*#__PURE__*/React.createElement(Button, {
      onClick: function onClick() {
        return onNavigate('login');
      },
      className: "px-8 py-3 text-lg",
      "data-name": "start-button"
    }, t.startButton))));
  } catch (error) {
    console.error('Home page error:', error);
    reportError(error);
    return null;
  }
}