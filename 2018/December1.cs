using System;
using System.Text.RegularExpressions;
using System.Collections.Generic;
public class December1 {
    public static int CalculateFrequency(String[] frequencies) {
        int frequency = 0;

        foreach (var freq in frequencies) {
            frequency += Int32.Parse(freq);
        }
        return frequency;
    }

    public static int firstRepeatedFrequency(String[] frequencies) {
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