

def read_to_list(file: str) -> list[float]:
    lines: list[float] = []
    with open(file, "r") as f:
        lines = [ line.strip() for line in f.readlines()]

    return lines
        
