import requests as req
import json, os

class api_handler():
    
    def __init__(self):
        self.api_url = os.environ.get("API_URL") 

    def call(self,url) -> json:
        headers = {} #? Add token when needed.
        response = req.request("GET", url, headers=headers)
        return response.json()

    def push_data(self,url:str) -> json:
        payload={}
        headers = {}
        response = req.request("POST", url, headers=headers, data=payload)
        return response.json()
    
    
    def buy(self,strategy_name:str, quantity:float) -> json:
        return self.push_data(f'{self.api_url}binance/futures/order/buy?strategy={strategy_name}&quantity={quantity}')

    def sell(self,strategy_name:str, quantity:float) -> json:
        return self.push_data(f'{self.api_url}binance/futures/order/sell?strategy={strategy_name}&quantity={quantity}')

    def currency_price(self,symbol:str) -> json:
        return self.call(f'{self.api_url}price?symbol={symbol}')

    def get_asset_infos(self, asset:str, account_name:str) -> json:
        return self.call(f'{self.api_url}/binance/futures/account?asset={asset}&account_name={account_name}')
    
    def hourly_data(self, symbol:str, x:int) -> json:
        return self.call(f'{self.api_url}data/hourly?symbol={symbol}&x={x}')

    def create_strategy(self, symbol:str, strategy_name:str) -> json:
        return self.call(f'{self.api_url}strategy/create?symbol={symbol}&name={strategy_name}')

    def get_strategies(self) -> json:
        return self.call(f'{self.api_url}strategy/get')

    def delete_strategy(self, strategy_name:str) -> json:
        return self.call(f'{self.api_url}strategy/delete?name={strategy_name}')

    def convert(self, currency_1:str, currency_2:str, quantity:float) -> json:
        return self.call(f'{self.api_url}convert?from={currency_1}&to={currency_2}&quantity={quantity}')
    
    def get_position(self, symbol:str, account_name:str) -> json:
        return self.call(f'{self.api_url}/binance/futures/position?symbol={symbol}&account_name={account_name}')
        
    def __del__(self):
        pass

    def __enter__(self):
        self.__init__()
        return self
    
    def __exit__(self, exc_type, exc_value, tb):
        if exc_type is not None:
            print(exc_type, exc_value, tb)
            return False
        return True