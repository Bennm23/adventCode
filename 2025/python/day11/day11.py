from functools import lru_cache
from lib import read_to_input


class Device:
    def __init__(self, line: str):
        self.name = line.split(":")[0]
        self.connections = line.split(":")[1][1:].strip().split()
    
    def is_out(self) -> bool:
        return self.connections[0] == "out"

devices = [ Device(line.strip()) for line in read_to_input("day11/input") ]
device_map = {}
for device in devices:
    device_map[device.name] = device

@lru_cache(maxsize=None)
def search(curr: Device, found_dac: bool, found_fft: bool) -> int:
    if curr.is_out():
        return 1 if found_dac and found_fft else 0
    else:
        if curr.name == "dac":
            found_dac = True
        if curr.name == "fft":
            found_fft = True

        cntr = 0
        for conn in curr.connections:
            cntr += search(device_map[conn], found_dac, found_fft)
        return cntr

def p1() -> int:
    return search(device_map["you"], True, True)

def p2() -> int:
    return search(device_map["svr"], False, False)

print("Part 1:", p1()) # 708
print("Part 2:", p2()) # 545394698933400