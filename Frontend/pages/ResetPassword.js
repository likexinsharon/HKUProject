function ResetPassword({ onNavigate }) {
    try {
        const [formData, setFormData] = React.useState({
            password: '',
            confirmPassword: ''
        });
        const [errors, setErrors] = React.useState({});
        const [success, setSuccess] = React.useState(false);

        const handleSubmit = async (e) => {
            e.preventDefault();
            const newErrors = {};

            if (!validatePassword(formData.password)) {
                newErrors.password = 'Password must be at least 8 characters';
            }
            if (formData.password !== formData.confirmPassword) {
                newErrors.confirmPassword = 'Passwords do not match';
            }

            setErrors(newErrors);

            if (Object.keys(newErrors).length === 0) {
                try {
                    // Handle password reset logic here
                    console.log('Password reset:', formData);
                    setSuccess(true);
                } catch (error) {
                    console.error('Password reset error:', error);
                }
            }
        };

        if (success) {
            return (
                <div className="auth-container flex items-center justify-center px-4" data-name="reset-password-success">
                    <div className="auth-card w-full max-w-md p-8 text-center">
                        <i className="fas fa-check-circle text-green-500 text-5xl mb-4"></i>
                        <h2 className="text-2xl font-bold mb-4">Password Reset Successful</h2>
                        <p className="text-gray-600 mb-4">
                            Your password has been successfully reset. You can now login with your new password.
                        </p>
                        <Button onClick={() => onNavigate('login')}>
                            Login
                        </Button>
                    </div>
                </div>
            );
        }

        return (
            <div className="auth-container flex items-center justify-center px-4" data-name="reset-password-container">
                <div className="auth-card w-full max-w-md p-8" data-name="reset-password-card">
                    <h1 className="form-title" data-name="reset-password-title">Reset Password</h1>
                    <form onSubmit={handleSubmit} data-name="reset-password-form">
                        <Input
                            type="password"
                            label="New Password"
                            value={formData.password}
                            onChange={(e) => setFormData({...formData, password: e.target.value})}
                            error={errors.password}
                            placeholder="Enter new password"
                        />
                        <Input
                            type="password"
                            label="Confirm New Password"
                            value={formData.confirmPassword}
                            onChange={(e) => setFormData({...formData, confirmPassword: e.target.value})}
                            error={errors.confirmPassword}
                            placeholder="Confirm new password"
                        />
                        <Button type="submit">Reset Password</Button>
                    </form>
                </div>
            </div>
        );
    } catch (error) {
        console.error('Reset password page error:', error);
        reportError(error);
        return null;
    }
}
