import time
from pathlib import Path

__NS_TO_S = 1e-9  #To Seconds
__NS_TO_MS = 1e-6 #To Milliseconds
__NS_TO_US = 1e-3 #To Microseconds

def filePath(fileName):
    if Path('/home/benn').is_dir():
        return '/home/benn/CODE/adventCode/' + fileName
    return '/home/bennmellinger/CODE/adventCode/' + fileName

    

def runAndPrintMicros(fn):
    __runWithTime(__NS_TO_US, 'Microseconds', fn)

def runAndPrintMillis(fn):
    __runWithTime(__NS_TO_MS, 'Milliseconds', fn)

def runAndPrintSeconds(fn):
    __runWithTime(__NS_TO_S, 'Seconds', fn)

def __runWithTime(transition, type, fn):
    start = time.perf_counter_ns() * transition
    fn()
    end = time.perf_counter_ns() * transition
    print('Duration ', type, ' = ', end - start)
    

    
val = "rn=1"
v2 = "cd-"

print("one = ", val.find("="))