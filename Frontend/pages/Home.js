function Home({ onNavigate }) {
    try {
        const { language } = React.useContext(LanguageContext);
        const t = translations[language].home;

        const features = [
            { icon: 'fa-chart-line', text: t.features.marketData },
            { icon: 'fa-money-bill-trend-up', text: t.features.trading },
            { icon: 'fa-wallet', text: t.features.management },
            { icon: 'fa-brain', text: t.features.ai }
        ];

        return (
            <div className="hero-container" data-name="home-container">
                <div className="hero-overlay"></div>
                <div className="hero-content h-full">
                    <nav className="flex justify-end p-6 relative z-10" data-name="home-nav">
                        <div className="flex items-center space-x-6">
                            <LanguageSelector />
                            <button
                                onClick={() => onNavigate('about')}
                                className="text-white hover:text-blue-300 transition-colors"
                                data-name="about-us-button"
                            >
                                {t.aboutUs}
                            </button>
                        </div>
                    </nav>
                    
                    <div className="flex flex-col items-center justify-center h-[calc(100%-80px)] px-4 text-center" data-name="hero-content">
                        <h1 className="text-4xl md:text-6xl font-bold text-white mb-4" data-name="hero-title">
                            {t.title}
                        </h1>
                        <p className="text-xl md:text-2xl text-white mb-8" data-name="hero-subtitle">
                            {t.subtitle}
                        </p>
                        
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-12" data-name="features-grid">
                            {features.map((feature, index) => (
                                <div key={index} className="feature-item" data-name={`feature-${index}`}>
                                    <i className={`fas ${feature.icon} feature-icon`}></i>
                                    <span>{feature.text}</span>
                                </div>
                            ))}
                        </div>
                        
                        <Button
                            onClick={() => onNavigate('login')}
                            className="px-8 py-3 text-lg"
                            data-name="start-button"
                        >
                            {t.startButton}
                        </Button>
                    </div>
                </div>
            </div>
        );
    } catch (error) {
        console.error('Home page error:', error);
        reportError(error);
        return null;
    }
}
