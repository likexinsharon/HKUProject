function LanguageSelector() {
    try {
        const [isOpen, setIsOpen] = React.useState(false);
        const { language, setLanguage } = React.useContext(LanguageContext);

        const languages = {
            en: 'English',
            zh: '中文'
        };

        return (
            <div className="language-selector" data-name="language-selector">
                <div
                    className="flex items-center cursor-pointer"
                    onClick={() => setIsOpen(!isOpen)}
                    data-name="language-selector-trigger"
                >
                    <i className="fas fa-globe mr-2"></i>
                    <span>{languages[language]}</span>
                </div>
                {isOpen && (
                    <div className="language-menu" data-name="language-menu">
                        {Object.entries(languages).map(([code, name]) => (
                            <div
                                key={code}
                                className={`px-4 py-2 hover:bg-gray-100 cursor-pointer ${
                                    language === code ? 'text-blue-600' : 'text-gray-700'
                                }`}
                                onClick={() => {
                                    setLanguage(code);
                                    setIsOpen(false);
                                }}
                                data-name={`language-option-${code}`}
                            >
                                {name}
                            </div>
                        ))}
                    </div>
                )}
            </div>
        );
    } catch (error) {
        console.error('LanguageSelector component error:', error);
        reportError(error);
        return null;
    }
}
