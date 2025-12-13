from functools import lru_cache
from lib import read_to_input


class Device:
    def __init__(self, line: str):
        self.name = line.split(":")[0]
        self.connections = line.split(":")[1][1:].strip().split(" ")
    
    def is_out(self) -> bool:
        return self.connections[0] == "out"

devices = [ Device(line.strip()) for line in read_to_input("day11/input") ]
device_map = {}
for device in devices:
    device_map[device.name] = device

def p1() -> int:
    total_paths = 0

    start = device_map["you"]

    visited = set()

    def search(curr:Device, path: str) -> int:
        if curr.is_out():
            return 1
        else:
            cntr = 0
            for conn in curr.connections:
                new_path = path + conn
                if new_path not in visited:
                    visited.add(new_path)
                    cntr += search(device_map[conn], new_path)
                    # visited.remove(new_path)
            return cntr

    for conn in start.connections:
        curr = "you" + conn
        paths = search(device_map[conn], curr)
        total_paths += paths


    return total_paths

# @lru_cache(maxsize=None)
def connected_to(target: str, all: set[str], visited: set[str]) -> None:
    if target not in device_map:
        return

    if target in visited:
        return
    visited.add(target)

    if target in device_map:
        for conn in device_map[target].connections:
            if conn not in all:
                all.add(conn)
                connected_to(conn, all, visited)
    # device = device_map[target]
    for device in devices:
        if target in device.connections:
            all.add(device.name)
            connected_to(device.name, all, visited)

def p2() -> int:
    total_paths = 0

    start = device_map["svr"]

    visited = set()

    @lru_cache(maxsize=None)
    def search(curr: Device, path: str) -> int:
        if curr.is_out() and "dac" in path and "fft" in path:
            return 1
        else:
            cntr = 0
            for conn in curr.connections:
                new_path = path + "," + conn
                if new_path not in visited:
                    visited.add(new_path)
                    if conn == "out":
                        continue
                    cntr += search(device_map[conn], new_path)
                    # visited.remove(new_path)
            return cntr

    for conn in start.connections:
        print("Processing connection:", conn)
        curr = "svr" + "," + conn
        paths = search(device_map[conn], curr)
        total_paths += paths


    return total_paths

# print("Part 1:", p1()) # 708
print("Part 2:", p2()) # 