import ctypes
import os
import time
import requests
import pycurl
from StringIO import StringIO 

def main():
    path = os.path.join(os.path.abspath("."), "ibm6418.so")
    lib = ctypes.CDLL(path)

    #lib.localServer() [ERROR -- TERMINAL FREEZES]
    print "GOLANG REQUEST"

    url = raw_input("ENTER URL: ")
    method = raw_input("ENTER METHOD: ")

    user_nURL = ctypes.c_char_p(url)
    user_nMETHOD = ctypes.c_char_p(method)
    
    start = time.time()
    for i in range(100):
        print "Request: ", i+1, " - "
        req = lib.makeRequestTest(user_nMETHOD.value, user_nURL.value) #[ERROR STORING VALUE] 
        print "\n"
    
    stop = time.time()

    totalTimeg = stop - start
    print "\nTOTAL TIME ELASPED: ", totalTimeg

    print "This is REQ: ", req # PRINTS REQUEST[ERROR]
    time.sleep(3)
    #-------------------------------------------------------------------------------------------------------------------------------------
    print "\nPYTHON REQUEST"
    print "URL: ", url

    start = time.time()
    for i in range(100):
        req = requests.get(url)
        print "Request: ", i+1, " - ", req.text, "\n"
    stop = time.time()
    
    totalTimep = stop - start
    print "TOTAL TIME ELASPED: ", totalTimep
    
    time.sleep(3)
    #-------------------------------------------------------------------------------------------------------------------------------------
    print "\nPYTHON-pycURL REQUEST"
    print "URL: ", url

    buffer = StringIO()

    start = time.time()
    for i in range(100):
        c = pycurl.Curl()
        c.setopt(c.URL, url)
        c.setopt(c.WRITEDATA, buffer)
        c.perform()
        body = buffer.getvalue()
        print "Request: ", i+1, " - ", body, "\n"
        c.close()
    stop = time.time()

    totalTimepc = stop - start
    time.sleep(3)
    #-------------------------------------------------------------------------------------------------------------------------------------
    print "GOLANG TIME: ", totalTimeg
    print "PYTHON TIME: ", totalTimep
    print "PYTHON-pycURL TIME: ", totalTimepc

    print "\nProgram has ended...\n"
    #-------------------------------------------------------------------------------------------------------------------------------------
if __name__ == "__main__":
    main()