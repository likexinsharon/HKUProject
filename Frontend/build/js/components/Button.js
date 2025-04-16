"use strict";

function Button(_ref) {
  var type = _ref.type,
    onClick = _ref.onClick,
    children = _ref.children,
    disabled = _ref.disabled;
  try {
    return /*#__PURE__*/React.createElement("button", {
      type: type || 'button',
      onClick: onClick,
      disabled: disabled,
      className: "w-full py-2 px-4 rounded-lg text-white font-medium transition-colors ".concat(disabled ? 'bg-gray-400 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700'),
      "data-name": "button"
    }, children);
  } catch (error) {
    console.error('Button component error:', error);
    reportError(error);
    return null;
  }
}