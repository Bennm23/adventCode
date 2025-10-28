
def read_to_list(file: str) -> list[str]:
    lines: list[str] = []
    with open(file, "r") as f:
        lines = [ line.strip() for line in f.readlines()]
    return lines
        

def read_to_int_list(file: str) -> list[int]:
    return [int(line) for line in read_to_list(file)]
