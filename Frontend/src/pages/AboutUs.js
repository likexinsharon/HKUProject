function AboutUs({ onNavigate }) {
    try {
        const { language } = React.useContext(LanguageContext);
        
        const team = [
            {
                name: language === 'en' ? 'Li Kexin' : '李可馨',
                role: language === 'en' ? 'Backend Developer' : '后端开发',
                image: 'Pictures/member1.jpg'
            },
            {
                name: language === 'en' ? 'Xie Yujian' : '谢钰鉴',
                role: language === 'en' ? 'Backend Developer' : '后端开发',
                image: 'Pictures/member2.jpg'
            },
            {
                name: language === 'en' ? 'Li Haitao' : '李海涛',
                role: language === 'en' ? 'Frontend Developer' : '前端开发',
                image: 'Pictures/member3.jpg'
            },
            {
                name: language === 'en' ? 'Wang Ziqi' : '汪子淇',
                role: language === 'en' ? 'Frontend Developer' : '前端开发',
                image: 'Pictures/member4.jpg'
            },
            {
                name: language === 'en' ? 'Li Zihan' : '李子寒',
                role: language === 'en' ? 'AI Model Engineer' : 'AI模型工程师',
                image: 'Pictures/member5.jpg'
            }
        ];

        return (
            <div className="about-page" data-name="about-us-container">
                <div className="about-content max-w-6xl mx-auto">
                    <div className="nav-container flex justify-between items-center">
                        <h1 className="text-4xl font-bold text-white">
                            {language === 'en' ? 'About Us' : '关于我们'}
                        </h1>
                        <div className="flex items-center space-x-4">
                            <div className="text-white">
                                <LanguageSelector />
                            </div>
                            <button
                                onClick={() => onNavigate('home')}
                                className="text-white hover:text-blue-300 transition-colors"
                                data-name="back-to-home"
                            >
                                <i className="fas fa-arrow-left mr-2"></i>
                                {language === 'en' ? 'Back to Home' : '返回首页'}
                            </button>
                        </div>
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8" data-name="team-grid">
                        {team.map((member, index) => (
                            <div
                                key={index}
                                className="team-member-card"
                                data-name={`team-member-${index}`}
                            >
                                <img
                                    src={member.image}
                                    alt={member.name}
                                    className="w-32 h-32 rounded-full mx-auto mb-4 object-cover border-4 border-white shadow-md"
                                />
                                <h3 className="text-xl font-semibold mb-2 text-gray-800">{member.name}</h3>
                                <p className="text-gray-600">{member.role}</p>
                            </div>
                        ))}
                    </div>

                    <div className="about-card" data-name="mission-section">
                        <h2 className="text-2xl font-semibold mb-4 text-gray-800">
                            {language === 'en' ? 'Our Mission' : '我们的使命'}
                        </h2>
                        <p className="text-gray-600 leading-relaxed">
                            {language === 'en' 
                                ? 'We are dedicated to making cryptocurrency trading accessible and efficient for everyone. Our platform combines cutting-edge technology with user-friendly interfaces to provide the best trading experience.'
                                : '我们致力于让加密货币交易对每个人都变得简单易行。我们的平台将尖端技术与用户友好的界面相结合，提供最佳的交易体验。'
                            }
                        </p>
                    </div>
                </div>
            </div>
        );
    } catch (error) {
        console.error('AboutUs page error:', error);
        reportError(error);
        return null;
    }
}
