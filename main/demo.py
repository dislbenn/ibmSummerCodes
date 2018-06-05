import ctypes
import os
import time
import requests
import grequests
import gohttp

def main():
    path = os.path.join(os.path.abspath("."), "demo.so")
    lib = ctypes.CDLL(path)
    #-------------------------------------------------------------------------------------------------------------------------------------
    print "GO REQUEST" # BEGIN GOLANG GET REQUEST CALL
    # lib.localServer() # SETUP LOCALSERVER FROM GO [ERROR] - FREEZES TERMINAL
    url = raw_input("ENTER URL: ")

    #Make Request
    for i in range(1000): # RETURN REQUEST 100 TIMES
        print "Request: ", i+1 # Prints the request number
        #lib.getRequest(url)# Make Request # ERROR
    
    start = time.time()
    #for i in range(1000):
        #req = lib.getRequest(url) # [ERROR]
        #print(req)
    #   print 
   
    #stop = time.time()
    #totalTime = stop - start
   #print "Golang Request - Time Elasped: ", totalTime # TOTAL TIME ELASPED FOR GOLANG REQUEST CALL
    #-------------------------------------------------------------------------------------------------------------------------------------
    time.sleep(2) 
    #-------------------------------------------------------------------------------------------------------------------------------------
    print "PYTHON REQUEST" # BEGIN PYTHON GET REQUEST CALL
    req = requests.get(url)
   
    print "Request status_code: ", req.status_code
    time.sleep(3)

    start = time.time()
    for i in range(100): # RETURN REQUEST 100 TIMES
        req = requests.get(url) # Make Request
        print "Request: ", i+1, req.text #Prints the request

    stop  = time.time()
    totalTime = stop - start
    print "Python Request - Time Elasped: ", totalTime # TOTAL TIME ELASPED FOR PYTHON REQUEST CALL
    #-------------------------------------------------------------------------------------------------------------------------------------
if __name__ == "__main__":
    main()