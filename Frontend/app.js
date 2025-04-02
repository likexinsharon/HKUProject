function App() {
    try {
        const [currentPage, setCurrentPage] = React.useState('home');
        const [language, setLanguage] = React.useState('en');

        const handleNavigate = (page) => {
            setCurrentPage(page);
        };

        const renderPage = () => {
            switch (currentPage) {
                case 'home':
                    return <Home onNavigate={handleNavigate} />;
                case 'about':
                    return <AboutUs onNavigate={handleNavigate} />;
                case 'login':
                    return <Login onNavigate={handleNavigate} />;
                case 'register':
                    return <Register onNavigate={handleNavigate} />;
                case 'forgot-password':
                    return <ForgotPassword onNavigate={handleNavigate} />;
                case 'reset-password':
                    return <ResetPassword onNavigate={handleNavigate} />;
                default:
                    return <Home onNavigate={handleNavigate} />;
            }
        };

        return (
            <LanguageContext.Provider value={{ language, setLanguage }}>
                <div className="min-h-screen bg-gray-100" data-name="app-container">
                    {renderPage()}
                </div>
            </LanguageContext.Provider>
        );
    } catch (error) {
        console.error('App component error:', error);
        reportError(error);
        return null;
    }
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);
