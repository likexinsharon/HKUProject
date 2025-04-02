function Input({ type, label, value, onChange, error, placeholder }) {
    try {
        return (
            <div className="form-group" data-name="input-group">
                <label className="form-label" data-name="input-label">{label}</label>
                <input
                    type={type}
                    value={value}
                    onChange={onChange}
                    placeholder={placeholder}
                    className={`w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 ${
                        error ? 'border-red-500' : 'border-gray-300'
                    }`}
                    data-name="input-field"
                />
                {error && <p className="error-message" data-name="input-error">{error}</p>}
            </div>
        );
    } catch (error) {
        console.error('Input component error:', error);
        reportError(error);
        return null;
    }
}
