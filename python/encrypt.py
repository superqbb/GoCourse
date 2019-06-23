import sys
import json
import mysql.connector

# 要先修改uic中的phone,email长度，以确保足够的加密字符串长度
# 需要先安装mysql-connector 
# python -m pip install mysql-connector

def encryptData(str):
    return str+"123"

config = {
    'host': '192.168.188.150',
    'user': 'devuser',
    'password': '',
    'port': 3306,
    'database': 'uic',
    'charset': 'utf8'
}
print(sys.path)
try:
    cnn = mysql.connector.connect(**config)
except mysql.connector.Error as e:
    print('[ERROR] connect fails!{}'.format(e))

cursor = cnn.cursor()
try:
    sql_count = 'select count(1) from user'
    cursor.execute(sql_count)
    count = cursor.fetchone()[0]
    print('[INFO] Encrypting data, total count：',count)

    sql_query = 'select id,name,cnname,email,phone from user'
    cursor.execute(sql_query)
    faild_list = []
    for id, name, cnname, email, phone in cursor.fetchall():
        try:
            encrypt_email = encryptData(email)
            encrypt_phone = encryptData(phone)
            sql_update = "update user set email= %s,phone= %s where id = %s"
            val = (encrypt_email, encrypt_phone, id)
            cursor.execute(sql_update, val)
            cnn.commit()
        except Exception as e:
            faild_list.append({'id':id,'name':name,'cnname':cnname,'email':email,'phone':phone})
            print("[ERROR] update uic faild ,user id=%s,error=%s"%(id,e))
    f = open("faild_list.json",'w')
    f.write(json.dumps(faild_list))
    f.close()
    print('[INFO] Data encrypt complete, total:%s, faild:%i'%(count, len(faild_list)))
except mysql.connector.Error as e:
        print('[ERROR] query error!{}'.format(e))
finally:
    cursor.close()
    cnn.close()



