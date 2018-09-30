#!/usr/bin/env python3
# -*- coding: UTF-8 -*-
from module.mysqlOJ import MysqlOJ
from module.fileOJ import FileOJ
from websocket_server import WebsocketServer
import bcrypt


class WebSocketOJ:
    def __init__(self, mysqloj, fileoj, serveroj):
        self.mysql = mysqloj
        self.fileoj = fileoj
        self.serveroj = serveroj

    def new_client(self, client, server):
        print("New client connected and was given id %d" % client['id'])
        client['check_login'] = False
        client['send_login'] = True
        client['send_exam'] = False
        client['send_qp'] = False
        client['send_code'] = False
        client['send_code_submit'] = ""

    def client_left(self, client, server):
        print("Client(%d) disconnected" % (client['id']))

    def message_received(self, client, server, message):
        if client['check_login']:
            self.__check_login(client, server, message)
        # elif client['send_exam']:
        #     self.__exam(client, server, message)
        elif client['send_qp']:
            self.__qp(client, server, message)
        elif client['send_code']:
            self.__code(client, server, message)

        if message == "login":
            client['check_login'] = True
        elif message == "exam":
            client['send_exam'] = True
        elif message == "qp":
            client['send_qp'] = True
        elif message == "code":
            client['send_code'] = True

    def __check_login(self, client, server, message):
        if client['send_login']:
            results = mysqloj.get_login(message.rstrip())
            if len(results):
                for data in results:
                    client['user_id'] = data[0]
                    client['user_name'] = data[1]
                    client['user_email'] = data[2]
                    client['user_passwd_hash'] = data[3]
                client['send_login'] = False
            else:
                var_check = bcrypt.hashpw(
                    bytes("F", encoding="utf8"), bcrypt.gensalt())
                server.send_message(client, bytes("login", encoding="utf8"))
                server.send_message(client, var_check)
                print("%s(client_id->%d)=>Email does not exist" %
                      (client['address'], client['id']))
                client['check_login'] = False
        else:
            if bcrypt.checkpw(bytes(message.rstrip(), encoding="utf8"), bytes(client['user_passwd_hash'], encoding="utf8")):
                var_check = bcrypt.hashpw(
                    bytes("T", encoding="utf8"), bcrypt.gensalt())
                server.send_message(client, bytes("login", encoding="utf8"))
                server.send_message(client, var_check)
                print("[%s](client_id->%d:user_id->%s)=>Sign in suceesfully" %
                      (client['user_email'], client['id'], client['user_id']))
            else:
                var_check = bcrypt.hashpw(
                    bytes("F", encoding="utf8"), bcrypt.gensalt())
                server.send_message(client, bytes("login", encoding="utf8"))
                server.send_message(client, var_check)
                print("[%s](client_id->%d:user_id->%s)=>Wrong password" %
                      (client['user_email'], client['id'], client['user_id']))
            client['send_login'] = True
            client['check_login'] = False

    # def __exam(self, client, server, message):
        # results = mysqloj.get_exam()
        # server.send_message(client, bytes(str(len(results)), encoding="utf8"))
        # if len(results):
        #     for data in results:
        #         for i in data:
        #             server.send_message(client, bytes(i, encoding="utf8"))
        # else:
        #     pass

    def __qp(self, client, server, message):
        results_QP = mysqloj.get_all_QP()
        server.send_message(client, bytes("qp", encoding="utf8"))
        server.send_message(client, bytes(
            str(len(results_QP)), encoding="utf8"))

        for data in results_QP:

            for index, i in enumerate(data):
                if index == 0:
                    i = str(i)
                    results_Ex = mysqloj.get_example(i)
                    server.send_message(client, bytes(
                        str(len(results_Ex)), encoding="utf8"))

                server.send_message(client, bytes(
                    i.replace("\r", ""), encoding="utf8"))

            for data in results_Ex:
                for i in data:
                    server.send_message(client, bytes(
                        i.replace("\r", ""), encoding="utf8"))

        client['send_qp'] = False

    def __code(self, client, server, message):
        if client['send_code_submit'] == "run":
            print("(client_id->%d:user_id->%s:qpid->%s)=>Receive data" %
                  (client['id'], client['user_id'], message))
            client['qpid'] = message
            client['send_code_submit'] = "run2"
        elif client['send_code_submit'] == "run2":
            file_save = open(
                'user_code/'+str(client['user_id'])+'_'+client['qpid']+'.py', 'wb')
            file_save.write(message.encode())
            # fp.write(str.encode(message))
            file_save.close()
            print("(client_id->%d:user_id->%s:qpid->%s)=>Compile" %
                  (client['id'], client['user_id'], client['qpid']))

            results = mysqloj.get_testing(client['qpid'])
            var_type = mysqloj.get_inputtype(client['qpid'])
            input_type = fileoj.transform_type(var_type[0][0])
            client['filepy'] = fileoj.add_suffix(
                str(client['user_id']), client['qpid'], input_type)
            answer_flag = True
            for data in results:
                output = fileoj.compile(client['filepy'], data[0])
                if output != data[1]:
                    server.send_message(client, bytes("code", encoding="utf8"))
                    server.send_message(
                        client, bytes(data[0], encoding="utf8"))
                    server.send_message(client, bytes(output, encoding="utf8"))
                    answer_flag = False
                    break

            if answer_flag:
                server.send_message(client, bytes("code", encoding="utf8"))
                server.send_message(client, bytes("T", encoding="utf8"))

            fileoj.del_file(client['filepy'])
            print("(client_id->%d:user_id->%s:qpid->%s)=>Compile complete" %
                  (client['id'], str(client['user_id']), client['qpid']))

            client['send_code_submit'] = ""
            client['send_code'] = False

        if message == "run":
            client['send_code_submit'] = "run"

        # server.send_message(client, bytes("hello", encoding="utf8"))
        # print("Client(%d) said: %s" % (client['id'], message))


if __name__ == '__main__':
    localhost = "mysql"
    user = "user"
    userPassword = "user123"
    db = "weboj"
    mysqloj = MysqlOJ(localhost, user, userPassword, db)
    if mysqloj.mysqloj_connect():
        fileoj = FileOJ()
        server = WebsocketServer(1121, host='172.17.0.4')
        wsoj = WebSocketOJ(mysqloj, fileoj, server)
        server.set_fn_new_client(wsoj.new_client)
        server.set_fn_client_left(wsoj.client_left)
        server.set_fn_message_received(wsoj.message_received)
        server.run_forever()
