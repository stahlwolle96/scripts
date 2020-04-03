/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
 
 v2.43-1
 
    DESCRIPTION
 1. The script can send fleets from more than 1 planet/moon
 2. Check/Get EXPO Debris(if you are Discoverer)
 3. You can start this script at specific time
*/
homes = ["M:1:2:3"] // Replace M:1:2:3 with your coordinate - M for the moon, P for planet.
// You can add as many planets/moons you want - the home list must look like this: homes = ["M:1:2:3", "M:2:2:3"]

shipsList = {LARGECARGO: 3000, LIGHTFIGHTER: 10000, DESTROYER: 25, PATHFINDER: 0}/* Your can change ENTIRE List, even to left only 1 type of ships! 
If you set 0 to some type of the ships, the script will send ALL ships of this type at once!
IMPORTANT!!! This script accept the ships list literally and NOT calculate your ships depense of the free slots, so if you want to send more than 1 fleet per planet/moon, you must calculate very precious your ships before set the ships list!
*/

minusCurrentSystem = 3 // Set this as start destination of range coordinates - minus your current world's system
plusCurrentSystem = 5 // Set this as end destination of range coordinates - plus your current world's system

DurationOfExpedition = 1 // Set duration (in hours) of the EXPEDITION: minimum 1 - maximum 8
PathfindersDebris = true // Do you want to get EXPO debrises? true = YES / false = NO
Pnbr = 5  // The script will ignore debris less than for PATHFINDERS that you set - The Maximum PATHFINDERS is limited only of your PATHFINDERS on the current moon/planet! You can set this value from 1, to the number you want
PathfinderSystemsRange = true // Do you want to check/get EXPO debris in range systems? true = YES / false = NO
SystemsRange = false // Do you want to send your EXPO fleet to Range coordinates? true = YES / false = NO
Repeat = true // Do you want to repeat the full cycle of fleet sending? true = YES / false = NO
HowManyCycles = 5 // Set the limit of repeats of whole cycle of EXPO fleet sending - 0 means forewer

myTime = "09:33:00"// Set your start Time; Hour: 00 - 23, Minute: 00 - 59
useStartTime = false // Do you want to run this script at specific time every day? true = YES / false = NO

//-------
current = 0
err = nil
wrong = []
curentco = {}
homeworld = nil
i = 0
ei = 0
er = nil
flag = 0
cng = 0
cycle = 0
endFlag = 0
fleetFlag = 0
RepeatTimes = 1
if (Pnbr < 1) {Pnbr = 1}
for home in homes {
    for celestial in GetCachedCelestials() {
        if GetCachedCelestial(celestial) == GetCachedCelestial(homes[0]) {
            homeworld = GetCachedCelestial(homes[i])
            ei = ei + 1
        } else {flag = 1}
    }
    if flag == 1 {wrong += homes[i]}
    if i < len(homes)-1 {
        i++
        flag = 0
    }
}
if ei == len(homes) {homeworld = GetCachedCelestial(homes[0])}
if !IsDiscoverer() {
    Print("You are not Discoverer and cannot get the EXPO Debris!")
    PathfindersDebris = false
}
if useStartTime == false {
    hour, minute, sec = Clock()
    startHour = hour
    startMin = minute
    startSec = sec + 3
    if startSec >= 60 {
        startSec = startSec - 60
        startMin = startMin + 1
        if startMin >= 60 {
            startMin = startMin - 60
            startHour = startHour + 1
        }
        if startHour >= 24 {startHour = startHour - 24}
    }
    myTime = ""+startSec+" "+startMin+" "+startHour+" * * *"
}
if HowManyCycles == 0 {HowManyCycles = false}
if homeworld != nil {
    totalUsl = GetSlots().Total - GetFleetSlotsReserved()
    totalExpSlots = GetSlots().ExpTotal
    CronExec(myTime, func() {
        for home = current; home <= len(homes)-1; home++ {
            pp = 0
            Dtarget = 0
            homeworld = GetCachedCelestial(homes[home])
            fromSystem = homeworld.GetCoordinate().System - minusCurrentSystem
            if fromSystem < 1 {fromSystem = 1}
            toSystem = homeworld.GetCoordinate().System + plusCurrentSystem
            if fromSystem > 499 {toSystem = 499}
            crdn = fromSystem
            if homeworld.Coordinate.IsMoon() {
                Print("Your Moon is: "+homeworld.Coordinate)
            } else {Print("Your Planet is: "+homeworld.Coordinate)}
            times = totalExpSlots
            currentTime = 0
            if SystemsRange == true && cycle >= len(homes)-1 {
                for id, num in curentco {
                    if id == homes[home] {crdn = num}
                }
            }
            totalSlots = totalUsl
            if PathfindersDebris == true {
                dflag = 0
                abr = 0
                nbr = 0
                curSystem = fromSystem
                if PathfinderSystemsRange == false {
                    curSystem = homeworld.GetCoordinate().System
                    toSystem = homeworld.GetCoordinate().System
                }
                for system = curSystem; system <= toSystem; system++ {
                    myShips, _ = homeworld.GetShips()
                    systemInfos, _ = GalaxyInfos(homeworld.GetCoordinate().Galaxy, system)
                    Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+system+":"+16)
                    Debris, _ = ParseCoord("D:"+homeworld.GetCoordinate().Galaxy+":"+system+":"+16)
                    slots = GetSlots().InUse
                    Sleep(Random(1500, 3700))
                    if slots < totalSlots {
                        Print("Checking "+Dtarget)
                        if systemInfos.ExpeditionDebris.PathfindersNeeded >= Pnbr {
                            pp = systemInfos.ExpeditionDebris.PathfindersNeeded
                            if systemInfos.ExpeditionDebris.Metal == 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                            if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal == 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal)}
                            if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal+" and Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                            fleet, _ = GetFleets()
                            for f in fleet {
                                if f.Mission == RECYCLEDEBRISFIELD && f.ReturnFlight == false {
                                    if Debris == f.Destination {
                                        if f.Ships.Pathfinder < pp {
                                            abr = pp - f.Ships.Pathfinder
                                        } else {dflag = 1}
                                    }
                                }
                            }
                            if dflag == 0 {
                                f = NewFleet()
                                f.SetOrigin(homeworld)
                                f.SetDestination(Dtarget)
                                f.SetSpeed(HUNDRED_PERCENT)
                                f.SetMission(RECYCLEDEBRISFIELD)
                                if abr == 0 {
                                    nbr = systemInfos.ExpeditionDebris.PathfindersNeeded
                                } else {nbr = abr}
                                if nbr > myShips.Pathfinder {nbr = myShips.Pathfinder}
                                f.AddShips(PATHFINDER, nbr)
                                a, b = f.SendNow()
                                if b == nil {
                                    if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                                    if nbr > 1 {
                                        Print(nbr+" Pathfinders are sended successfully!")
                                    } else {Print(nbr+" Pathfinder is sended successfully!")}
                                } else {
                                    if nbr > 1 {
                                        Print("The Pathfinders are NOT sended! "+b)
                                    } else {Print("The Pathfinder is NOT sended! "+b)}
                                    break
                                }
                            } else {Print("Needed ships already are sended!")}
                        }
                    }
                }
                if pp == 0 {Print("Not found any debris!")}
            }
            slots = GetSlots().InUse
            Sleep(Random(1000, 3000))
            if slots < totalSlots {
                slots = GetSlots().ExpInUse
                totalSlots = totalExpSlots
            }
            Sleep(Random(1000, 3000))
            if err != nil {slots = totalSlots}
            if slots < totalSlots {
                for time = currentTime; time < times; time++ {
                    myShips, _ = homeworld.GetShips()
                    tt = 0
                    rtt = 0
                    ExpFleet = {}
                    totalSlots = totalUsl
                    slots = GetSlots().InUse
                    Sleep(Random(1500, 3000))
                    if slots < totalSlots {
                        slots = GetSlots().ExpInUse
                        totalSlots = totalExpSlots
                    }
                    if slots < totalSlots {
                        if SystemsRange == false {
                            Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+homeworld.GetCoordinate().System+":"+16)
                        }
                        if SystemsRange == true {
                            if crdn > toSystem {crdn = fromSystem}
                            Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+crdn+":"+16)
                        }
                        explist = []
                        Sleep(Random(13000, 18000)) // For avoiding ban
                        fleet = NewFleet()
                        fleet.SetOrigin(homeworld)
                        fleet.SetDestination(Dtarget)
                        fleet.SetSpeed(HUNDRED_PERCENT)
                        fleet.SetMission(EXPEDITION)
                        sltPerWorld = times - time
                        if sltPerWorld == 0 {sltPerWorld = 1}
                        if len(shipsList) > 0 {
                            for ShipID, num in shipsList {
                                rtt = rtt + 1
                                if myShips.ByID(ShipID) != 0 {
                                    if num == 0 {
                                        ExpFleet[ShipID] = Floor(myShips.ByID(ShipID)/sltPerWorld)
                                        tt = tt + 1
                                    } else {
                                        if ShipID != PATHFINDER {
                                            if myShips.ByID(ShipID) >= num {
                                                ExpFleet[ShipID] = num
                                                tt = tt + 1
                                            }
                                        } else {
                                            if myShips.ByID(ShipID) < num {num = myShips.ByID(ShipID)}
                                            ExpFleet[ShipID] = num
                                            tt = tt + 1
                                        }
                                    }
                                }
                            }
                        }
                        fleet.SetDuration(DurationOfExpedition)
                        if rtt == tt {
                            for ShipID, nbr in ExpFleet {
                                fleet.AddShips(ShipID, nbr)
                                explist += ShipID+": "+nbr
                            }
                        }
                        a, err = fleet.SendNow()
                        if err == nil {
                            Print(explist+" are sended successfully to "+Dtarget)
                            if SystemsRange == true {
                                if crdn <= toSystem {crdn++}
                                curentco[homes[home]] = crdn
                            }
                        } else {
                            time = times
                            Print("The fleet is NOT sended! "+err)
                            er = err
                            err = nil
                        }
                        if home >= len(homes)-1 {err = er}
                    }
                    slots = GetSlots().ExpInUse
                    if err != nil {slots = totalSlots}
                }
            }
            if cycle <= len(homes)-1 {cycle++}
            if home >= len(homes)-1 {
                for slots == totalSlots {
                    delay = Random(7*60, 12*60) // 7 - 12 minutes in seconds
                    if Repeat == true {
                        slots = GetSlots().ExpInUse
                        expslots = slots
                        if err != nil {
                            if slots != 0 {
                                for slots == expslots {
                                    delay = Random(7*60, 12*60) // 7 - 12 minutes in seconds
                                    if err.Error() == "no ships to send" {
                                        Print("Please wait till ships lands! Recheck after "+ShortDur(delay))
                                    } else {Print("Will recheck after "+ShortDur(delay))}
                                    Sleep(delay*1000)
                                    expslots = GetSlots().ExpInUse
                                    if slots > expslots {
                                        err = nil
                                        er = nil
                                    }
                                    delay = nil
                                }
                            } else {
                                if cng == 0 {
                                    Print("All your EXPO ships are on the ground! Please, check your deuterium and make sure that you set the ships list correctly, then start the script again!")
                                    RepeatTimes = HowManyCycles
                                    useStartTime = false
                                    endFlag = 1
                                }
                            }
                        } else {
                            Print("All slots are busy now! Please, wait "+ShortDur(delay))
                            Sleep(delay*1000)
                            slots = GetSlots().ExpInUse
                        }
                    } else {
                        slots = 1
                        totalSlots = 3
                    }
                    delay = nil
                }
                if RepeatTimes != HowManyCycles {
                    if HowManyCycles != false {
                        if Repeat == true {Print("You make full cycle of fleet sending "+RepeatTimes+"!")}
                        RepeatTimes++
                    }
                    current = -1
                    if Repeat == true {home = current}
                } else {
                    if endFlag == 0 {Print("You have reached the limit of repeats that you have set")}
                    Sleep(3000)
                    
                }
            }
            Sleep(Random(1000, 3000))
        }
        if useStartTime == false {StopScript(__FILE__)}
    })
} else {
    Print("You typed wrong coordinates! - "+wrong)
    StopScript(__FILE__)
}
<-OnQuitCh
