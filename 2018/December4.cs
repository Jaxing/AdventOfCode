using System;
using System.Collections.Generic;
using System.Text.RegularExpressions;

public enum EventType {
    Sleep=1,Wake=2,Start=3

}
public class Event : IComparable<Event> {
    public int guardId;
    public EventType type;
    public DateTime timestamp;
    public Event(int guardId, string timestamp, EventType type) {
        this.guardId = guardId;
        this.timestamp = DateTime.Parse(timestamp);
        this.type = type;
    }

    public int CompareTo(Event e) {
        return timestamp.CompareTo(e.timestamp);
    }
}
public class Guard {
    public int id;
    public List<Event> events = new List<Event>();
    public Guard(int id) {
        this.id = id;
    }

    public int sleepDuration() {
        this.events.Sort();
        int timeAsleep = 0;

        for (int i=0; i < this.events.Count - 1; i++) {
            if (this.events[i].type == EventType.Sleep) {
                //Assume that there are always a awake after asleep
                timeAsleep += this.events[i+1].timestamp.Subtract(this.events[i].timestamp).Minutes;
            }
        }

        return timeAsleep;
    }

    public (int,int) mostAsleepAt() {
        this.events.Sort();
        Dictionary<int,int> sleepFrequency = new Dictionary<int, int>();

        for (int i=0; i < this.events.Count - 1; i++) {
            if (this.events[i].type == EventType.Sleep) {
                DateTime current = this.events[i].timestamp;
                var start = DateTime.Now;
                while (current < this.events[i+1].timestamp) {
                    sleepFrequency[current.Minute] = sleepFrequency.GetValueOrDefault(current.Minute) + 1;
                    current = current.AddMinutes(1);
                }
            }
        }

        int mostSleep = 0;
        int sleepyMinute = 0;

        foreach (var (k,v) in sleepFrequency) {
            if (v > mostSleep) {
                mostSleep = v;
                sleepyMinute = k;
            }
        }

        return (sleepyMinute, mostSleep);
    }

    // override object.Equals
    public override bool Equals(object obj)
    {
        //
        // See the full list of guidelines at
        //   http://go.microsoft.com/fwlink/?LinkID=85237
        // and also the guidance for operator== at
        //   http://go.microsoft.com/fwlink/?LinkId=85238
        //
        
        if (obj == null || GetType() != obj.GetType())
        {
            return false;
        }

        return this.id == ((Guard) obj).id;
    }
    
    // override object.GetHashCode
    public override int GetHashCode()
    {
        return this.id * 97;
    }
}
public class December4 {

    public static List<Event> parse(string[] events) {
        List<Event> eventList = new List<Event>();

        foreach (var e in events) {
            var timestamp = e.Substring(1,16);
            var typeId = e[19].ToString().ToUpper();
            var idString = Regex.Match(e, @"#(?<id>\d*)").Groups["id"].ToString();
            var type = EventType.Wake;
            var id = 0;
            
            switch(typeId) {
                case "G":
                    id = Int32.Parse(idString);
                    type = EventType.Start;
                    break;
                case "W":
                    type = EventType.Wake;
                    break;
                case "F":
                    type = EventType.Sleep;
                    break;
            }

            eventList.Add(new Event(id, timestamp, type));
        }

        return eventList;
    }

    public static int sleepyNumber(string[] events) {
        List<Event> eventList = parse(events);
        List<Guard> guards = new List<Guard>();
        eventList.Sort();
        Guard currentGuard = new Guard(eventList[0].guardId);

        foreach (Event e in eventList) {
            switch(e.type) {
                case EventType.Sleep:
                    currentGuard.events.Add(e);
                    break;
                case EventType.Start:
                    if (guards.Exists(x => x.id ==e.guardId )) {
                        currentGuard = guards.Find(x => x.id == e.guardId);
                    } else {
                        currentGuard = new Guard(e.guardId);
                        guards.Add(currentGuard);
                    }
                    currentGuard.events.Add(e);
                    break;
                case EventType.Wake:
                    currentGuard.events.Add(e);
                    break;
            } 
        }

        int mostSleep = 0;
        Guard sleepyGuard = null;

        foreach (var guard in guards) {
            int guardSleep = guard.sleepDuration();
            if (guardSleep > mostSleep) {
                mostSleep = guardSleep;
                sleepyGuard = guard;
            }
        }

        var (a, b) = sleepyGuard.mostAsleepAt() ;

        
        return a * sleepyGuard.id;
    }

    public static int reliableSleeper(string[] events) {
         List<Event> eventList = parse(events);
        List<Guard> guards = new List<Guard>();
        eventList.Sort();
        Guard currentGuard = new Guard(eventList[0].guardId);

        foreach (Event e in eventList) {
            switch(e.type) {
                case EventType.Sleep:
                    currentGuard.events.Add(e);
                    break;
                case EventType.Start:
                    if (guards.Exists(x => x.id ==e.guardId )) {
                        currentGuard = guards.Find(x => x.id == e.guardId);
                    } else {
                        currentGuard = new Guard(e.guardId);
                        guards.Add(currentGuard);
                    }
                    currentGuard.events.Add(e);
                    break;
                case EventType.Wake:
                    currentGuard.events.Add(e);
                    break;
            } 
        }

        Guard goodSleeper = null;
        int mostSleepy = 0;
        int mostSleepyAt = 0;

        foreach (var guard in guards) {
            var (sleepMinute, timesSlept) = guard.mostAsleepAt();

            if (timesSlept > mostSleepy) {
                mostSleepy = timesSlept;
                mostSleepyAt = sleepMinute;
                goodSleeper = guard;
            }
        }

        return goodSleeper.id * mostSleepyAt;
    }

}