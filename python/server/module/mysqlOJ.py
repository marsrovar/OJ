# -*- coding: UTF-8 -*-
import pymysql
import sys


class MysqlOJ:
    def __init__(self, localhost, user, userPassword, db):
        self.host = localhost
        self.user = user
        self.userPW = userPassword
        self.db = db

    def mysqloj_connect(self):
        try:
            pymysql.connect(self.host, self.user, self.userPW, self.db)
            return True
        except:
            info = sys.exc_info()
            print(info[0], ":", info[1])

    def get_login(self, email):
        try:
            db = pymysql.connect(self.host, self.user, self.userPW, self.db)
            cursor = db.cursor()
            sql = "select id, name, email, password from users where email="+"\'"+email+"\'"
            cursor.execute(sql)
            results = cursor.fetchall()
            db.close()
            return results
        except:
            info = sys.exc_info()
            print(info[0], ":", info[1])

    def get_exam(self):
        try:
            db = pymysql.connect(self.host, self.user, self.userPW, self.db)
            cursor = db.cursor()
            sql = "select id, num, title, content from "
            cursor.execute(sql)
            results = cursor.fetchall()
            db.close()
            return results
        except:
            info = sys.exc_info()
            print(info[0], ":", info[1])

    def get_all_QP(self):
        try:
            db = pymysql.connect(self.host, self.user, self.userPW, self.db)
            cursor = db.cursor()
            sql = "select id, title, content, code_function from questionpools"
            cursor.execute(sql)
            results = cursor.fetchall()
            db.close()
            return results
        except:
            info = sys.exc_info()
            print(info[0], ":", info[1])

    def get_example(self, qpid: str):
        try:
            db = pymysql.connect(self.host, self.user, self.userPW, self.db)
            cursor = db.cursor()
            sql = "select example from questionpools_examples where qpid="+qpid
            cursor.execute(sql)
            results = cursor.fetchall()
            db.close()
            return results
        except:
            info = sys.exc_info()
            print(info[0], ":", info[1])

    def get_inputtype(self, qpid: str):
        try:
            db = pymysql.connect(self.host, self.user, self.userPW, self.db)
            cursor = db.cursor()
            sql = "select inputtype from questionpools where id="+qpid
            cursor.execute(sql)
            results = cursor.fetchall()
            db.close()
            return results
        except:
            info = sys.exc_info()
            print(info[0], ":", info[1])

    def get_testing(self, qpid: str):
        try:
            db = pymysql.connect(self.host, self.user, self.userPW, self.db)
            cursor = db.cursor()
            sql = "select input, output from questionpools_testings where qpid="+qpid
            cursor.execute(sql)
            results = cursor.fetchall()
            db.close()
            return results
        except:
            info = sys.exc_info()
            print(info[0], ":", info[1])


if __name__ == '__main__':
    pass
