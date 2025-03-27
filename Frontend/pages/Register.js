function Register({ onNavigate }) {
    try {
        const { language } = React.useContext(LanguageContext);
        const t = translations[language].auth;

        const [formData, setFormData] = React.useState({
            username: '',
            email: '',
            password: '',
            confirmPassword: ''
        });
        const [errors, setErrors] = React.useState({});

        const handleSubmit = async (e) => {
            e.preventDefault();
            const newErrors = {};

            if (!validateUsername(formData.username)) {
                newErrors.username = 'Username must be at least 3 characters';
            }
            if (!validateEmail(formData.email)) {
                newErrors.email = 'Please enter a valid email address';
            }
            if (!validatePassword(formData.password)) {
                newErrors.password = 'Password must be at least 8 characters';
            }
            if (formData.password !== formData.confirmPassword) {
                newErrors.confirmPassword = 'Passwords do not match';
            }

            setErrors(newErrors);

            if (Object.keys(newErrors).length === 0) {
                try {
                    // Handle registration logic here
                    console.log('Registration attempt:', formData);
                } catch (error) {
                    console.error('Registration error:', error);
                }
            }
        };

        return (
            <div className="auth-page" data-name="register-container">
                <div className="auth-card" data-name="register-card">
                    <div className="flex justify-end mb-4">
                        <LanguageSelector />
                    </div>
                    <h1 className="form-title" data-name="register-title">{t.register}</h1>
                    <form onSubmit={handleSubmit} data-name="register-form">
                        <Input
                            type="text"
                            label="Username"
                            value={formData.username}
                            onChange={(e) => setFormData({...formData, username: e.target.value})}
                            error={errors.username}
                            placeholder="Choose a username"
                        />
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
                        <Input
                            type="password"
                            label="Confirm Password"
                            value={formData.confirmPassword}
                            onChange={(e) => setFormData({...formData, confirmPassword: e.target.value})}
                            error={errors.confirmPassword}
                            placeholder="Confirm your password"
                        />
                        <Button type="submit">{t.register}</Button>
                        <p className="mt-4 text-center text-gray-600" data-name="login-prompt">
                            {t.alreadyHaveAccount}{' '}
                            <a
                                href="#"
                                onClick={(e) => {
                                    e.preventDefault();
                                    onNavigate('login');
                                }}
                                className="form-link"
                            >
                                {t.loginHere}
                            </a>
                        </p>
                    </form>
                </div>
            </div>
        );
    } catch (error) {
        console.error('Register page error:', error);
        reportError(error);
        return null;
    }
}
