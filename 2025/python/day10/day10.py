from functools import lru_cache
import itertools
import math
import re
from lib import read_to_input

Combination = tuple[list[int], ...]
Combinations = list[Combination]

class Machine:
    def __init__(self, line: str):
        self.indicators: list[bool] = [ c == "#" for c in line[1:].split("]")[0]]

        schematics_groups = re.compile(r"\((.*?)\)").findall(line)
        self.schematics: list[list[int]] = [ [ int(x) for x in group.split(",") ] for group in schematics_groups ]

        joltage = re.compile(r"\{(.*?)\}").findall(line)
        self.joltage: list[int] = [ int(x) for x in joltage[0].split(",") ]

    def __str__(self):
        return f"Machine:\n\tIndicators: {self.indicators}\n\tSchematics: {self.schematics}\n\tJoltage: {self.joltage}\n"

    def sequence_works(self, sequence: list[int]) -> bool:
        start = [False for _ in self.indicators]
        for press in sequence:
            for button in press:
                start[button] = not start[button]
        
        for i in range(len(start)):
            if start[i] != self.indicators[i]:
                return False
        return True

    def valid_for_joltage(self, sequence: list[int]) -> bool:
        button_press_count: dict[int, int] = {}
        for press in sequence:
            for button in press:
                if button not in button_press_count:
                    button_press_count[button] = 0
                button_press_count[button] += 1
        
        for i in range(len(self.joltage)):
            count = button_press_count.get(i, 0)
            if count != self.joltage[i]:
                return False
        return True

    def build_potential_sequences(self) -> list[Combinations]:
        all_sequences: list[Combinations] = []

        # Build sequences
        for i, counter in enumerate(self.joltage):
            # The schematics that affect this indicator
            these_schematics: list[list[int]] = []
            for schematic in self.schematics:
                if i in schematic:
                    these_schematics.append(schematic)

            # Get all combinations of these schematics that achieve the joltage required
            combinations: Combinations = list(itertools.combinations_with_replacement(these_schematics, r=counter))
            all_sequences.append(combinations)
        
        return all_sequences

    # @lru_cache(maxsize=None)
    def build_sequence_for_joltage(self, button_index: int, joltage: list[int]) -> Combinations:
        # Try to build a sequence that achieves the given joltage
        # for the given index
        joltage_goal = joltage[button_index]
        these_schematics: list[list[int]] = []
        for schematic in self.schematics:
            if button_index in schematic:
                these_schematics.append(schematic)

        combinations = list(itertools.combinations_with_replacement(these_schematics, r=joltage_goal))
        return combinations

machines = [ Machine(line.strip()) for line in read_to_input("day10/input") ]

def p1() -> int:
    count = 0

    for machine in machines:
        print("Processing Machine")
        found = False
        curr_len = 1
        while not found:
            perm_iter = itertools.permutations(machine.schematics, r=curr_len)
            for perm in perm_iter:
                if machine.sequence_works(perm):
                    print("Found valid sequence of length", curr_len)
                    found = True
                    break
            if not found:
                curr_len += 1

        count += curr_len
    return count

def apply_sequence(machine: Machine, button_press: Combination, curr_joltage: list[int], best_found: int, curr_length: int, curr_index: int) -> int:
    curr_length += len(button_press)
    # If this sequence is already longer than best, abort
    if curr_length >= best_found:
        return best_found

    # print("Applying sequence:", button_press, "to joltage:", curr_joltage,  " curr_length:", curr_length, " curr_index:", curr_index, " best_found:", best_found)
    best_len = best_found

    new_joltage = curr_joltage.copy()
    for press in button_press:
        for button in press:
            new_joltage[button] -= 1

    if all( j == 0 for j in new_joltage ):
        return curr_length
    elif any( j < 0 for j in new_joltage ):
        # If this resulted in an invalid joltage
        return best_found

    curr_index += 1

    if curr_index >= len(new_joltage):
        return best_found


    new_opts = machine.build_sequence_for_joltage(curr_index, new_joltage)
    # No combinations are valid to solve this joltage sequence
    if len(new_opts) == 0:
        return best_found

    for opt in new_opts:
        total_len = apply_sequence(machine, opt, new_joltage, best_found, curr_length, curr_index)
        if total_len < best_len:
            best_len = total_len
        if total_len < best_found:
            best_found = total_len

    return best_len

def p2() -> int:
    count = 0

    for machine in machines:
        print("Processing Machine")
        
        # min_presses = max(machine.joltage)
        # print("Minimum presses needed:", min_presses)

        curr_joltage = machine.joltage.copy()
        best_found = 10000

        options = machine.build_sequence_for_joltage(0, curr_joltage)
        # print(f"Indicator {0} has {len(options)} options to achieve joltage {curr_joltage[0]}")
        # print(options)

        for opt in options:
            # print(opt)
            total_len = apply_sequence(machine, opt, curr_joltage, best_found, 0, 0)
            if total_len < best_found:
                best_found = total_len
        
        count += best_found
        # For each joltage index, build combos of schematics that achieve that joltage

        # List all potential sequences that achieve the joltage requirements
        # for each indicator I, indexed by the indicator number
        # sequences: list[Combinations] = machine.build_potential_sequences()

        # all_combos = itertools.product(*sequences)
        # for combo in all_combos:
        #     print(combo)
        # for index, combinations in enumerate(sequences):
        #     # Each combinations list contains all the potential combinations of schematics that achieve the joltage
        #     # Each successive index mus
        #     print(f"Indicator {index}: {len(combinations)} potential combinations")
        #     for seq in combinations:
        #         print(f"   {seq}")


    return count

# print("P1:", p1()) #545
print("P2:", p2()) #


