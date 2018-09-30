# -*- coding: UTF-8 -*-
import subprocess
import os
import sys


class FileOJ:
    path = os.path.abspath(os.path.dirname(os.path.dirname(__file__)))

    def add_suffix(self, userid: str, qpid: str, input_type):
        try:
            output = subprocess.check_output("grep def "+self.path+"/user_code/"+userid +
                                             "_"+qpid+".py |cut -d'(' -f 1|awk 'BEGIN{ORS=\"\"}{print $2}'", shell=True)
            fo = open(self.path+"/user_code/"+userid+"_"+qpid+".py", "a")
            str1 = "\nimport sys\nvar=Solution()\nvar2="+input_type + \
                "(sys.argv[1])\nprint(var." + \
                output.decode('utf-8')+"(var2), end='')\n"
            fo.write(str1)
            fo.close()
            return userid+"_"+qpid+".py"
        except subprocess.CalledProcessError:
            print('Exception handled')

    def compile(self, filename: str, input_data: str):
        output = subprocess.Popen("timeout 5 python "+self.path+"/user_code/"+filename+" "+input_data +
                                  ";var=$(echo $?);if [ $var != 0 ] && [ $var != 1 ];then echo 'TimeOutError'; fi", shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)

        output_str = ""
        while output.poll() is None:
            line = output.stdout.readline()
            if line:
                output_str += line.decode('utf-8')
        if output.returncode == 0:
            return output_str
        else:
            return 'failed'

    def transform_type(self, input_type):
        if input_type == 'int':
            return "int"
        elif input_type == 'string':
            return "str"

    def del_file(self, filename: str):
        os.remove(self.path+"/user_code/"+filename)


if __name__ == '__main__':
    pass
