"use strict";

var translations = {
  en: {
    home: {
      title: "All-in-One Cryptocurrency Trading & Analysis Platform!",
      subtitle: "Helping You stay Ahead of the Market with Ease!!",
      features: {
        marketData: "Real-time Market Data",
        trading: "Simulated Trading",
        management: "Asset Management",
        ai: "AI Evaluation"
      },
      startButton: "Let's start!",
      aboutUs: "About Us"
    },
    auth: {
      login: "Login",
      register: "Register",
      email: "Email Address",
      password: "Password",
      username: "Username",
      confirmPassword: "Confirm Password",
      chooseUsername: "Choose a username",
      confirmYourPassword: "Confirm your password",
      forgotPassword: "Forgot Password?",
      rememberMe: "Remember me",
      dontHaveAccount: "Don't have an account?",
      registerHere: "Register here",
      alreadyHaveAccount: "Already have an account?",
      loginHere: "Login here"
    }
  },
  zh: {
    home: {
      title: "一站式加密货币交易和分析平台！",
      subtitle: "轻松帮您领先市场！",
      features: {
        marketData: "实时市场数据",
        trading: "模拟交易",
        management: "资产管理",
        ai: "人工智能评估"
      },
      startButton: "开始使用！",
      aboutUs: "关于我们"
    },
    auth: {
      login: "登录",
      register: "注册",
      email: "电子邮箱",
      password: "密码",
      username: "用户名",
      confirmPassword: "确认密码",
      chooseUsername: "选择一个用户名",
      confirmYourPassword: "确认您的密码",
      forgotPassword: "忘记密码？",
      rememberMe: "记住我",
      dontHaveAccount: "还没有账号？",
      registerHere: "在此注册",
      alreadyHaveAccount: "已有账号？",
      loginHere: "在此登录"
    }
  }
};
var LanguageContext = React.createContext({
  language: 'en',
  setLanguage: function setLanguage() {}
});