using System;
using System.Text.RegularExpressions;
using System.Collections.Generic;
public class December1 {
    static void Main(string[] args) {
        if (args[1] == "1") {
            string[] input = Regex.Split(args[2], ",");
            Console.WriteLine(CalculateFrequency(input));
        } else if (args[1] == "2") {
            string[] input = Regex.Split(args[2], ",");
            Console.WriteLine(firstRepeatedFrequency(input));
        }
    }

    static int CalculateFrequency(String[] frequencies) {
        int frequency = 0;

        foreach (var freq in frequencies) {
            frequency += Int32.Parse(freq);
        }
        return frequency;
    }

    static int firstRepeatedFrequency(String[] frequencies) {
        int frequency = 0;
        HashSet<int> encounteredFrequencies = new HashSet<int>();
        int i = 0;

        while (!encounteredFrequencies.Contains(frequency)) {
            encounteredFrequencies.Add(frequency);
            frequency += Int32.Parse(frequencies[i]);
            i++;
            if (i >= frequencies.Length) {
                i = 0;
            }
        }


        return frequency;
    }
}