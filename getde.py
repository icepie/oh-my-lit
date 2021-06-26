import requests
import re
import curlify

username = "B19071121"
password = "029753"

session = requests.Session()

headers = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36',
}

def get_sydl()->dict:
    response = session.get('http://zhyd.sec.lit.edu.cn/zhyd/sydl/index', headers=headers, verify=False)

    print(curlify.to_curl(response.request))

    sydlPage = response.text

    all_data = {}
    try:
        all_data['sp'] = re.search(r'剩余电量<span class="mui-badge mui-badge-success">(.*?)</span></li>', sydlPage).group(1)
        all_data['ra'] = re.search(r'剩余金额<span class="mui-badge mui-badge-success">(.*?)</span></li>', sydlPage).group(1)
        all_data['rs'] = re.search(r'剩余补助<span class="mui-badge mui-badge-success">(.*?)</span></li>', sydlPage).group(1)
        all_data['rsb'] = re.search(r'剩余补助金额<span class="mui-badge mui-badge-success">(.*?)</span></li>', sydlPage).group(1)
    except:
        send_login(username,password)
        pass

    print(session.cookies)

    return all_data

def send_login(username:str,password:str):
    response = session.get('http://ids.lit.edu.cn/authserver/login', headers=headers, verify=False, allow_redirects=False)

    print(curlify.to_curl(response.request))

    loginPage = response.text

    lt = re.search(r'name="lt" value="(.*?)"', loginPage).group(1)
    execution = re.search(r'name="execution" value="(.*?)"', loginPage).group(1)
    eventId = re.search(r'name="_eventId" value="(.*?)"', loginPage).group(1)
    rmShown  = re.search(r'name="rmShown" value="(.*?)"', loginPage).group(1)

    data = {
        'username': username,
        'password': password,
        'lt': lt,
        'execution': execution,
        '_eventId': eventId,
        'rmShown': rmShown,
    }

    print(session.cookies)

    response = session.post('http://ids.lit.edu.cn/authserver/login', data=data, headers=headers, timeout=5, verify=False,allow_redirects=False)

    print(curlify.to_curl(response.request))




def login():
    send_login(username,password)

login()
print(get_sydl())