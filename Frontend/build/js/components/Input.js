"use strict";

function Input(_ref) {
  var type = _ref.type,
    label = _ref.label,
    value = _ref.value,
    onChange = _ref.onChange,
    error = _ref.error,
    placeholder = _ref.placeholder;
  try {
    return /*#__PURE__*/React.createElement("div", {
      className: "form-group",
      "data-name": "input-group"
    }, /*#__PURE__*/React.createElement("label", {
      className: "form-label",
      "data-name": "input-label"
    }, label), /*#__PURE__*/React.createElement("input", {
      type: type,
      value: value,
      onChange: onChange,
      placeholder: placeholder,
      className: "w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 ".concat(error ? 'border-red-500' : 'border-gray-300'),
      "data-name": "input-field"
    }), error && /*#__PURE__*/React.createElement("p", {
      className: "error-message",
      "data-name": "input-error"
    }, error));
  } catch (error) {
    console.error('Input component error:', error);
    reportError(error);
    return null;
  }
}