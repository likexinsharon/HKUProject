function Login({ onNavigate }) {
    try {
        const { language } = React.useContext(LanguageContext);
        const t = translations[language].auth;

        const [formData, setFormData] = React.useState({
            email: '',
            password: '',
            rememberMe: false
        });
        const [errors, setErrors] = React.useState({});

        const handleSubmit = async (e) => {
            e.preventDefault();
            const newErrors = {};

            if (!validateEmail(formData.email)) {
                newErrors.email = 'Please enter a valid email address';
            }
            if (!validatePassword(formData.password)) {
                newErrors.password = 'Password must be at least 8 characters';
            }

            setErrors(newErrors);

            if (Object.keys(newErrors).length === 0) {
                try {
                    // Handle login logic here
                    console.log('Login attempt:', formData);
                } catch (error) {
                    console.error('Login error:', error);
                }
            }
        };

        return (
            <div className="auth-page" data-name="login-container">
                <div className="auth-card" data-name="login-card">
                    <div className="flex justify-between items-center mb-4">
                        <button
                            onClick={() => onNavigate('home')}
                            className="text-blue-600 hover:text-blue-700 transition-colors flex items-center"
                            data-name="back-to-home"
                        >
                            <i className="fas fa-arrow-left mr-2"></i>
                            {language === 'en' ? 'Back to Home' : '返回首页'}
                        </button>
                        <div className="login-language-selector">
                            <LanguageSelector />
                        </div>
                    </div>
                    <h1 className="form-title" data-name="login-title">{t.login}</h1>
                    <form onSubmit={handleSubmit} data-name="login-form">
                        <Input
                            type="email"
                            label={t.email}
                            value={formData.email}
                            onChange={(e) => setFormData({...formData, email: e.target.value})}
                            error={errors.email}
                            placeholder={t.email}
                        />
                        <Input
                            type="password"
                            label={t.password}
                            value={formData.password}
                            onChange={(e) => setFormData({...formData, password: e.target.value})}
                            error={errors.password}
                            placeholder={t.password}
                        />
                        <div className="flex justify-between items-center mb-4" data-name="login-options">
                            <label className="flex items-center text-sm text-gray-600">
                                <input
                                    type="checkbox"
                                    checked={formData.rememberMe}
                                    onChange={(e) => setFormData({...formData, rememberMe: e.target.checked})}
                                    className="mr-2"
                                />
                                {t.rememberMe}
                            </label>
                            <a
                                href="#"
                                onClick={(e) => {
                                    e.preventDefault();
                                    onNavigate('forgot-password');
                                }}
                                className="form-link"
                            >
                                {t.forgotPassword}
                            </a>
                        </div>
                        <Button type="submit">{t.login}</Button>
                        <p className="mt-4 text-center text-gray-600" data-name="register-prompt">
                            {t.dontHaveAccount}{' '}
                            <a
                                href="#"
                                onClick={(e) => {
                                    e.preventDefault();
                                    onNavigate('register');
                                }}
                                className="form-link"
                            >
                                {t.registerHere}
                            </a>
                        </p>
                    </form>
                </div>
            </div>
        );
    } catch (error) {
        console.error('Login page error:', error);
        reportError(error);
        return null;
    }
}
