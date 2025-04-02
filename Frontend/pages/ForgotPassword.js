function ForgotPassword({ onNavigate }) {
    try {
        const [email, setEmail] = React.useState('');
        const [error, setError] = React.useState('');
        const [submitted, setSubmitted] = React.useState(false);

        const handleSubmit = async (e) => {
            e.preventDefault();
            if (!validateEmail(email)) {
                setError('Please enter a valid email address');
                return;
            }

            try {
                // Handle password reset request logic here
                console.log('Password reset requested for:', email);
                setSubmitted(true);
            } catch (error) {
                console.error('Password reset request error:', error);
                setError('Failed to send reset email. Please try again.');
            }
        };

        if (submitted) {
            return (
                <div className="auth-container flex items-center justify-center px-4" data-name="forgot-password-success">
                    <div className="auth-card w-full max-w-md p-8 text-center">
                        <i className="fas fa-check-circle text-green-500 text-5xl mb-4"></i>
                        <h2 className="text-2xl font-bold mb-4">Check Your Email</h2>
                        <p className="text-gray-600 mb-4">
                            We've sent password reset instructions to your email address.
                        </p>
                        <Button onClick={() => onNavigate('login')}>
                            Return to Login
                        </Button>
                    </div>
                </div>
            );
        }

        return (
            <div className="auth-container flex items-center justify-center px-4" data-name="forgot-password-container">
                <div className="auth-card w-full max-w-md p-8" data-name="forgot-password-card">
                    <h1 className="form-title" data-name="forgot-password-title">Forgot Password</h1>
                    <p className="text-gray-600 mb-6 text-center" data-name="forgot-password-description">
                        Enter your email address and we'll send you instructions to reset your password.
                    </p>
                    <form onSubmit={handleSubmit} data-name="forgot-password-form">
                        <Input
                            type="email"
                            label="Email Address"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            error={error}
                            placeholder="Enter your email"
                        />
                        <Button type="submit">Send Reset Link</Button>
                        <p className="mt-4 text-center text-gray-600" data-name="login-prompt">
                            Remember your password?{' '}
                            <a
                                href="#"
                                onClick={(e) => {
                                    e.preventDefault();
                                    onNavigate('login');
                                }}
                                className="form-link"
                            >
                                Login here
                            </a>
                        </p>
                    </form>
                </div>
            </div>
        );
    } catch (error) {
        console.error('Forgot password page error:', error);
        reportError(error);
        return null;
    }
}
