from api_handler import api_handler
from stable_baselines3 import PPO
from pprint import pprint
import numpy as np
from dotenv import load_dotenv
import json

def log_strategy(log_str):
    print(f'[ppo_supertrend] {log_str}')

actions_trad = {
    0:'SELL',
    1:'BUY'
}

# Load env
load_dotenv()
 
# Load and call the API
api = api_handler()
hourly_data = api.hourly_data("BTCUSDT", 10)

# Format the data for 
total_obs = [[1] if hour['supertrend'] == 'long' else [-1] for hour in hourly_data]
ar = np.array(total_obs)
newarr = ar.reshape(1,10,1)

# Load the model and predict
model = PPO.load('PPO_supertrend/best_model', env=None)
action, _states = model.predict(newarr)

log_strategy(f'Prediction action : {actions_trad[action[0]]}')

# Get Balance
infos = api.get_asset_infos('USDT','aztakur')
available_balance = int(float(infos['availableBalance']))

# Get Position
pos = api.get_position('BTCUSDT','aztakur')
positionAmount = float(pos['positionAmt'])

# If we have no position and bot want to buy
if positionAmount == 0.000 and action[0] == 1 : 
    quantity_to_buy = round(api.convert('USDT','BTC',available_balance*0.98)['quantity'],3)
    log_strategy(f"BUY : {api.buy('PPO_supertrend', quantity_to_buy)}")
elif positionAmount > 0.000 and action[0] == 0: # If we have a position ans want to sell it
    log_strategy(f"SELL : {api.sell('PPO_supertrend', positionAmount)}")
else:
    log_strategy("HOLD")
