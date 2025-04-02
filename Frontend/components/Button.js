function Button({ type, onClick, children, disabled }) {
    try {
        return (
            <button
                type={type || 'button'}
                onClick={onClick}
                disabled={disabled}
                className={`w-full py-2 px-4 rounded-lg text-white font-medium transition-colors ${
                    disabled
                        ? 'bg-gray-400 cursor-not-allowed'
                        : 'bg-blue-600 hover:bg-blue-700'
                }`}
                data-name="button"
            >
                {children}
            </button>
        );
    } catch (error) {
        console.error('Button component error:', error);
        reportError(error);
        return null;
    }
}
