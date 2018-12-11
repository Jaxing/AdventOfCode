using System;
using System.Collections.Generic;
using System.Text.RegularExpressions;
using System.Text;
using System.Threading;

public class Elf : IComparable {
    public int time = 0;
    static int idCounter = 0;
    public int id;

    public Elf() {
        this.id = idCounter++;
    }
    Instruction assignedInstruction;

    public void assignInstruction(Instruction ins) {
        this.time += ins.getId() - 4;
        this.assignedInstruction = ins;
    }

    public Instruction getAssigned() {
        return assignedInstruction;
    }
    public void doInstruction() {
        this.assignedInstruction.setDone();
    }

    public bool isAvailable() {
        return this.assignedInstruction == null || this.assignedInstruction.isDone();
    }

    public void wait(int time) {
        this.time += time;
    }

    public int CompareTo(object obj) {
        if (obj.GetType() != GetType()) {
            throw new ArgumentException();
        }
        Elf e = ((Elf) obj);
        int i = this.time.CompareTo(e.time);

        if (i == 0) {
            if (e.isAvailable() == this.isAvailable()) {
                return i;
            }
            if (e.isAvailable()) {
                return -1;
            }
            return 1;
        }

        return i;
    }

    public override string ToString() {
        return this.id.ToString();
    }
}
public class Instruction : IComparable {
    List<Instruction> prerequisits = new List<Instruction>();
    List<Instruction> dependent = new List<Instruction>();
    bool done = false;
    char id;

    int endTime;

    public Instruction(char id) {
        this.id = id;
    }

    public char getId() {
        return id;
    }
    public List<Instruction> getDependent() {
        return this.dependent;
    }
    public void addPrerequisit(Instruction ins) {
        this.prerequisits.Add(ins);
    }
    public void addDependent(Instruction ins) {
        this.dependent.Add(ins);
    } 
    public bool isDone() {
        return done;
    }

    public void setDone() {
        this.done = true;
    }

    public void updateEndTime(Instruction pre) {
        int idOffset = this.id - 64;
        this.endTime = Math.Max(this.endTime, pre.endTime + idOffset);
    }
    public bool canStart() {
        bool ready = true;
        
        foreach(var prerequisit in this.prerequisits) {
            ready = ready && prerequisit.done;
        }

        return ready;
    }

    public int CompareTo(object obj) {
        if ( obj.GetType() != GetType()) {
            throw new ArgumentException();
        }
        Instruction o = (Instruction) obj;
        return this.id.CompareTo(o.id);
    }

    public override string ToString() {
        return this.id.ToString();
    }
}

public class December7 {
    public static List<Instruction> parse(string[] instructions) {
        Dictionary<char, Instruction> nameToInstruction = new Dictionary<char, Instruction>();
        List<Instruction> instructionQueue = new List<Instruction>();
        Regex re = new Regex(@"Step (?<ins>[A-Z])(\k<ins>) can begin.");

        foreach (var instruction in instructions) {
            string[] parts = Regex.Split(instruction, " must be finished before step ");
            char thisInstruction = parts[0][parts[0].Length-1];
            char dependentInstruction = parts[1][0];
            Instruction current = new Instruction(thisInstruction);

            if (nameToInstruction.ContainsKey(thisInstruction)) {
                current = nameToInstruction[thisInstruction];
            } else {
                nameToInstruction[thisInstruction] = current;
            }

            if (nameToInstruction.ContainsKey(dependentInstruction)) {
                nameToInstruction[dependentInstruction].addPrerequisit(current);
                current.addDependent(nameToInstruction[dependentInstruction]);
            } else {
                Instruction dependent = new Instruction(dependentInstruction);
                dependent.addPrerequisit(current);
                current.addDependent(dependent);
                nameToInstruction[dependentInstruction] = dependent;
            }
        }

        foreach (var key in nameToInstruction.Keys){
            if (nameToInstruction[key].canStart()) {
                instructionQueue.Add(nameToInstruction[key]);
            }
        }

        return instructionQueue;
    }

    public static void updateQueue(Instruction current, List<Instruction> instructionQueue) {
        instructionQueue.Remove(current);
        foreach (var dep in current.getDependent()){
            if (dep.canStart()) {
                instructionQueue.Add(dep);
            }
        }
    }

    public static void updateQueue(Instruction current, SortedSet<Instruction> instructionQueue) {
        
        foreach (var dep in current.getDependent()){
            if (dep.canStart()) {
                instructionQueue.Add(dep);
            }
        }
    }

    public static string instructionOrder(string[] instructions) {
        StringBuilder sb = new StringBuilder();
        var instructionQueue = parse(instructions);

        Console.WriteLine(instructionQueue.Count);
        while (instructionQueue.Count > 0) {
            instructionQueue.Sort();
            Instruction current = instructionQueue[0];
            current.setDone();
            instructionQueue.Remove(current);
            sb.Append(current.getId());
            updateQueue(current, instructionQueue);
        }

        return sb.ToString();
    }

    public static Elf getFirstUnassigned(List<Elf> elfs) {
        foreach(var elf in elfs) {
            if (elf.isAvailable()) {
                return elf;
            }
        }
        return null;
    }
    public static int timeToFinnish(string[] instructions) {
        StringBuilder sb = new StringBuilder();
        var instructionQueue = parse(instructions);

        List<Elf> elfs = new List<Elf>();
        for (int i=0; i < 5; i++) {
            var e = new Elf();
            elfs.Add(e);
        }

        int time = 0;
        bool someoneBusy = false;
        while (instructionQueue.Count > 0 || someoneBusy) {
            someoneBusy = false;
            instructionQueue.Sort();
            List<Instruction> toRemove = new List<Instruction>();

            foreach (var ins in instructionQueue) {
                elfs.Sort();
                var canditade = getFirstUnassigned(elfs);
                if (canditade != null && !toRemove.Contains(ins)) {
                    Console.WriteLine("Assigne {0} to {1} at {2}", ins, canditade, time);
                    canditade.assignInstruction(ins);
                    toRemove.Add(ins);
                    someoneBusy = true;
                }
            }
            time++;

            foreach (var ins in toRemove) {
                instructionQueue.Remove(ins);
            }

            foreach (var elf in elfs) {
                if (!elf.isAvailable()) {
                    if (elf.time == time) {
                        Console.WriteLine("{0} does {1} until {2}", elf, elf.getAssigned(), elf.time);
                        elf.doInstruction();
                        updateQueue(elf.getAssigned(), instructionQueue);
                        sb.Append(elf.getAssigned());
                    }
                } else {
                    elf.wait(1);
                }
            }

            foreach (var elf in elfs) {
                if (!elf.isAvailable()) {
                    someoneBusy = true;
                }
            }
        }

        Console.WriteLine(sb.ToString());

        int last = 0;
        Console.WriteLine(elfs.Count);

        elfs.Sort();
        foreach (var elf in elfs) {
            Console.WriteLine("{0} has local time {1}", elf, elf.time);
            last = Math.Max(last, elf.time);
        }

        return last;
    }
}