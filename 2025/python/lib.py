

def read_to_input(file: str = "input"):
    sample = []
    with open(file) as f:
        sample = [line.strip() for line in f.readlines()]
    return sample