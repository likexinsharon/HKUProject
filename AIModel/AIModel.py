import requests
import json
from openai import OpenAI

class AIAPIClient:
    def __init__(self):
        """
        初始化 AI API 客户端

        Args:
            ai_analysis_url (str): AI 分析接口的 URL
        """
        self.client = OpenAI(api_key="sk-04fae40c9c8c491ea37b3517711d19e4", base_url="https://api.deepseek.com")
        self.prompt_analysis = """
        你是一个专业的量化交易策略分析师。根据提供的信息，请分析并选择最适合的交易模型，并给出具体的参数建议和交易策略。

        请严格按照以下JSON格式输出，不要包含任何其他文字：

        {
            "selected_model": "选择的模型名称",
            "parameters_advice": {
                "recommended_parameters": {
                    // 根据选择的模型返回相应的参数对象
                },
                "reason": "推荐该参数设置的详细理由"
            },
            "trading_strategy": {
                "buy_range": "建议买入价格范围的具体描述",
                "sell_range": "建议卖出价格范围的具体描述"
            },
            "risk_management": {
                "take_profit": 止盈比例数值,
                "stop_loss": 止损比例数值,
                "risk_reason": "风险管理建议的详细说明"
            }
        }

        分析要求：
        1. 从可选模型中选择最适合当前市场情况的模型
        2. 为选择的模型推荐最优参数设置
        3. 基于模型特性给出具体的买卖价格范围建议
        4. 提供合理的止盈止损比例设置

        请基于以下输入信息进行分析：
        - 当前AI分析输出：{advice_information}
        - 可选模型：{available_models}
        - 模型参数选项：{model_parameters}
        """
        self.prompt_report = """
        你是一个专业的量化交易策略评估专家。请基于提供的交易数据、模型选择和参数设置，对交易策略进行全面评估。

        请严格按照以下JSON格式输出，不要包含任何其他文字：

        {
            "overall_evaluation": "对整体交易表现的综合评价",
            "model_suggestion": {
                "is_suitable": true/false,
                "alternative_models": ["替代模型1", "替代模型2"],
                "reason": "模型适用性分析和建议理由"
            },
            "parameter_evaluation": {
                "is_reasonable": true/false,
                "suggested_improvements": {
                    // 建议改进的参数对象
                },
                "reason": "参数评估和改进建议的理由"
            },
            "risk_evaluation": {
                "risk_exposure": "风险暴露程度的详细分析",
                "suggested_improvements": "风险管理改进的具体建议",
                "potential_risks": ["潜在风险1", "潜在风险2", "潜在风险3"]
            }
        }

        评估要求：
        1. 分析K线数据和交易点位的合理性
        2. 评估日收益率的稳定性和风险水平
        3. 判断所选模型是否适合当前市场特征
        4. 评估参数设置的有效性
        5. 识别潜在风险并提供改进建议

        请基于以下数据进行评估：
        - K线数据：{kline_data}
        - 交易点数据：{trade_points}
        - 日收益率：{daily_returns}
        - 使用模型：{selected_model}
        - 参数设置：{parameters_setting}
        - 历史分析记录：{history}
        """
        self.ai_analysis_url = ""
        self.selected_model = "None"
        self.model_parameters = {}
        self.parameters_advice = {}
        self.message_analysis = [{"role":"system","content":self.prompt_analysis}]
        self.message_report = [{"role":"system","content":self.prompt_report}]

    def call_ai_analysis(self, advice_information, available_models, model_available_parameters):
        """
        第一次调用 AI 分析接口
        """
        """
        调用 AI 分析接口

        Args:
            advice_information (str): 当前 AI 版本的分析输出
            available_models (list): 可选模型列表
            model_parameters (dict): 模型对应超参数字典
            history (list, optional): 之前的对话历史. Defaults to None.

        Returns:
            dict: AI 分析结果
        """
        payload = {
            "advice_information": advice_information,
            "available_models": available_models,
            "model_parameters": model_available_parameters
        }
        self.message_analysis.append({"role":"user","content":json.dumps(payload)})
        try:

            response = self.client.chat.completions.create(
                model="deepseek-reasoner",
                messages=self.message_analysis,
                response_format={
                    'type': 'json_object'
                }
            )
            # reasoning_content = response.choices[0].reasoning_content
            content = json.loads(response.choices[0].message.content)
            self.message_analysis.append(response.choices[0].message)

            # 保存当前对话到历史记录

            return content
        except requests.exceptions.RequestException as e:
            print(f"请求 AI 分析接口失败：{e}")
            return None
    def call_ai_report(self, kline_data, trade_points, daily_returns, selected_model, parameters_setting):
        """
        调用 AI报告窗口，输出：
        {
            "overall_evaluation": "string (整体评价)",
            "model_suggestion": {
                "is_suitable": "boolean (模型是否适⽤)",
                "alternative_models": "array (其他建议模型)",
                "reason": "string (建议理由)"
            },
            "parameter_evaluation": {
                "is_reasonable": "boolean (参数设置是否合理)",
                "suggested_improvements": "object (建议的参数改进)",
                "reason": "string (改进理由)"
            },
            "risk_evaluation": {
                "risk_exposure": "string (⻛险暴露分析)",
                "suggested_improvements": "string (⻛险管理改进建议)",
                "potential_risks": "array (识别的潜在⻛险)"
            }
}
        """
        payload = {
            "kline_data": kline_data,
            "trade_points": trade_points,
            "daily_returns": daily_returns,
            "selected_model": selected_model,
            "parameters_setting": parameters_setting,
            "history": str(self.message_analysis)
        }
        self.message_report.append({"role":"user","content":json.dumps(payload)})
        try:
            response = self.client.chat.completions.create(
                model="deepseek-reasoner",
                messages=self.message_report,
                response_format={
                    'type': 'json_object'
                }
            )
            self.message_report.append(response.choices[0].message)
            content = json.loads(response.choices[0].message.content)
            return content
        except requests.exceptions.RequestException as e:
            print(f"请求 AI 分析接口失败：{e}")
            return None

# 示例用法
if __name__ == "__main__":
    # 初始化 API 客户端（这里需要替换为你实际的 API URL）

    client = AIAPIClient()
    prompt_analysis = """
    你是一个专业的量化交易策略分析师。根据提供的信息，请分析并选择最适合的交易模型，并给出具体的参数建议和交易策略。

    请严格按照以下JSON格式输出，不要包含任何其他文字：

    {
        "selected_model": "选择的模型名称",
        "parameters_advice": {
            "recommended_parameters": {
                // 根据选择的模型返回相应的参数对象
            },
            "reason": "推荐该参数设置的详细理由"
        },
        "trading_strategy": {
            "buy_range": "建议买入价格范围的具体描述",
            "sell_range": "建议卖出价格范围的具体描述"
        },
        "risk_management": {
            "take_profit": 止盈比例数值,
            "stop_loss": 止损比例数值,
            "risk_reason": "风险管理建议的详细说明"
        }
    }

    分析要求：
    1. 从可选模型中选择最适合当前市场情况的模型
    2. 为选择的模型推荐最优参数设置
    3. 基于模型特性给出具体的买卖价格范围建议
    4. 提供合理的止盈止损比例设置

    请基于以下输入信息进行分析：
    - 当前AI分析输出：{advice_information}
    - 可选模型：{available_models}
    - 模型参数选项：{model_parameters}
    """
    prompt_report = """
    你是一个专业的量化交易策略评估专家。请基于提供的交易数据、模型选择和参数设置，对交易策略进行全面评估。

    请严格按照以下JSON格式输出，不要包含任何其他文字：

    {
        "overall_evaluation": "对整体交易表现的综合评价",
        "model_suggestion": {
            "is_suitable": true/false,
            "alternative_models": ["替代模型1", "替代模型2"],
            "reason": "模型适用性分析和建议理由"
        },
        "parameter_evaluation": {
            "is_reasonable": true/false,
            "suggested_improvements": {
                // 建议改进的参数对象
            },
            "reason": "参数评估和改进建议的理由"
        },
        "risk_evaluation": {
            "risk_exposure": "风险暴露程度的详细分析",
            "suggested_improvements": "风险管理改进的具体建议",
            "potential_risks": ["潜在风险1", "潜在风险2", "潜在风险3"]
        }
    }

    评估要求：
    1. 分析K线数据和交易点位的合理性
    2. 评估日收益率的稳定性和风险水平
    3. 判断所选模型是否适合当前市场特征
    4. 评估参数设置的有效性
    5. 识别潜在风险并提供改进建议

    请基于以下数据进行评估：
    - K线数据：{kline_data}
    - 交易点数据：{trade_points}
    - 日收益率：{daily_returns}
    - 使用模型：{selected_model}
    - 参数设置：{parameters_setting}
    - 历史分析记录：{history}
    """

    # 第一次调用 AI 分析接口的示例
    ai_advice_info = "This is the current AI analysis output."
    ai_available_models = ["Moving Average", "EMA Strategy", "RSI"]
    ai_model_params = {
        "Moving Average": {
            "short_window": 7,
            "long_window": 30,
            "symbol": "BTC",
            "initial_balance": 10000.0,
            "start_date": "2023-01-01",
            "end_date": "2023-12-31",

        },
        "EMA Strategy": {
            "short_period": 12,
            "long_period": 26,
            "symbol": "BTC",
            "initial_balance": 10000.0,
            "start_date": "2023-01-01",
            "end_date": "2023-12-31"
        },
        "RSI": {
            "rsi_window": 14,
            "overbought": 70.0,
            "oversold": 30.0,
            "symbol": "BTC",
            "initial_balance": 10000.0,
            "start_date": "2023-01-01",
            "end_date": "2023-12-31",

        }
    }

    ai_result = client.call_ai_analysis(ai_advice_info, ai_available_models, ai_model_params)
    print("AI 分析结果：", ai_result)
    kline_data = []
    trade_points = []
    daily_returns = []
    selected_model = ai_result['selected_model']
    parameters_setting = ai_result['parameters_advice']['recommended_parameters']
    # 第二次调用 AI 分析接口的示例，带上历史记录
    new_ai_advice_info = "This is another AI analysis output with history."
    ai_result_with_history = client.call_ai_report(kline_data, trade_points, daily_returns, selected_model, parameters_setting)
    print("AI 分析结果（带历史）：", ai_result_with_history)


    #####结果展示
    """
    AI 分析结果： {'selected_model': 'RSI', 'parameters_advice': {'recommended_parameters': {'rsi_window': 14, 'overbought': 70.0, 'oversold': 30.0}, 'reason': '推荐使用标准参数设置：RSI窗口为14天，超买阈值为70，超卖阈值为30。这些值基于历史数据和行业标准，在加密货币（如BTC）的波动市场中能有效识别超买超卖信号，提高交易信号的准确性。'}, 'trading_strategy': {'buy_range': '建议在RSI指标低于30时买入，表示超卖条件，适合在价格回调时入场。', 'sell_range': '建议在RSI指标高于70时卖出，表示超买条件，适合在价格高位时出场。'}, 'risk_management': {'take_profit': 0.1, 'stop_loss': 0.05, 'risk_reason': '设置止盈10%（0.10）以锁定利润，止损5%（0.05）以控制潜在损失。这种比例适用于比特币的高波动性，平衡风险与回报，防止过度损失并确保策略稳健。'}}
AI 分析结果（带历史）： {'overall_evaluation': '由于提供的K线数据、交易点数据和日收益率均为空数组，无法评估实际交易表现。基于所选RSI模型和标准参数，策略在理论上是合理的，适合波动市场，但缺乏数据验证，实际效果需通过回测确认。', 'model_suggestion': {'is_suitable': True, 'alternative_models': ['Moving Average', 'EMA Strategy'], 'reason': 'RSI模型适合当前市场特征（如加密货币的高波动性），能有效识别超买超卖信号。替代模型如移动平均适合趋势市场，EMA策略提供更灵敏的响应，可在不同市场条件下作为备选。'}, 'parameter_evaluation': {'is_reasonable': True, 'suggested_improvements': {}, 'reason': '参数设置（rsi_window=14, overbought=70.0, oversold=30.0）是行业标准，在历史回测中表现稳定，适用于大多数市场。由于缺乏数据，无法提出具体优化建议；建议通过回测测试不同窗口（如9或25）或阈值调整（如overbought至65）以适应特定波动。'}, 'risk_evaluation': {'risk_exposure': '策略在高波动市场中暴露于假信号风险，可能导致频繁交易和累积损失；日收益率缺失无法评估稳定性，但标准RSI在趋势反转时风险较高。', 'suggested_improvements': '实施严格的风险管理，如设置止盈10%和止损5%以控制损失，并结合仓位管理降低风险暴露。', 'potential_risks': ['市场反转风险', '过度交易风险', '流动性不足风险']}}

    """