# Overview of all systems

## Game Tick Order

The state of a game ticks forward in the following order

- colonist needs
- colonist conditions
- colonist happiness
- colonist actions

## Colonist

### Needs

A mutable state that is always in flex representing the needs and desires of a colonist. This system is the core driver
behind the colonist's utility-based decision making.

#### Example Needs

- Rest
- Food
- Water
- Security
- Belonging
- Fulfillment
- Power
- Joy
- Wealth
- Status

#### Notes

- basic needs decrease at a steady rate, some needs decrease with more activity, like you become more exhausted chopping wood than you would making clothes
- a colonist feels a greater sense of belonging the more they socialize with others and receive positive affirmation from others
- fear increases with disasters and danger, security increases with armor, protection, safe housing
- a colonist needs a job, title or something useful to do to feel fulfilled
- a colonist feels a need for status the more they interact with others who have status
- a colonist feels a desire for more wealth when exposed to differences in wealth, such as through trade, visiting other colonines, or interacting with peers who have wealth

A colonist's mood is high when their needs are fulfilled. A high mood leads to inspired work, discovering new ideas, inspiring others, respect from peers
A colonist's mood is low when their needs are unfulfilled. A low mood leads to mental breaks, permanent negative traits, violence and drug dependence

```go
type NeedsController interface {
  All() map[ColonistKey][]Need
  GetValue(ColonistKey, Need) Value
  Adjust(ColonistKey, Need, Value) (Value, err)
  Target(ColonistKey, Need, Value, Duration) (cancel func(), err)
  Ceiling(ColonistKey, Need, Value, Duration) (cancel func(), err)
  Floor(ColonistKey, Need, Value, Duration) (cancel func(), err)
}

type Need interface {
  // The base rate at which a need increases or decreases per tick
  GetBaseRate()
  // The rate at which a need increases or decreases per tick.
  // Factors in all modifiers to this need's rate and returns a final product of those modifers
  GetRate()
  // The default target toward which this need always trends. Typically this is zero
  GetBaseTargetValue()
  // The target toward which this need is trending.
  // Factors in all modifiers to this need's target and returns a final product of those targets.
  GetTargetValue()
  // Returns the current value of the need.
  GetValue()
  // Sets the value of the need
  SetValue()
  // Returns the name of the need as a string
  GetName()
  // Returns the floor of this need (the need cannot go any lower than this value)
  // If there is more than one floor, this method returns the highest floor.
  GetFloorValue()
  // Returns the ceiling of this need (the need cannot go any higher than this value)
  // // If there is more than one ceiling, this method returns the lowest ceiling.
  GetCeilingValue()
}
```

### Conditions

Temporary states on the coloinist that apply modifiers to their skills, needs or attributes.

#### Example Conditions

- Dazed
- Frightened
- Encumbered
- Exhausted
- Inspired
- Starving
- Thirsty
- Grieving
- Sick

#### Notes

Conditions can be applied when a need hits a certain threshold or through events and interactions with the world.
Conditions are inherently temporary and can be fulfilled by meeting a need, waiting for the condition to pass, to tending to the condition by a professional
If someone passed away, all the colonists that had a relationship with the departed would have "Grieving" set.

```go
type Conditions interface {
  Get(ColonistKey) []Condition
  Add(ColonistKey, Condition)
  Remove(ColonistKey, Condition)
}
```

### Happiness

### Traits

Permanent descriptors for a colonist that apply modifiers to their skills, needs or attributes.

Example Traits

- Great Memory
- Deeply Empathetic
- Superiority Complex

Example Colonist

```js
{
  colonists: [
    {
      key: "123123123123",
      name: "Ada Goodlove",
      age: 512,
      beauty: 5,
      strength: 7,
      dexterity: 10,
      charisma: 6,
      memory: 6,
      skills: {
        hunting: 6,
        crafting: 6,
        cooking: 6,
        construction: 6,
        gathering: 6,
        mining: 6,
        woodcutting: 6,
        science: 6,
        combat: 6,
        medicine: 6,
      },
      traits: [
        {
          type: "great_memory",
          description:
            "{colonist.name} has a great memory. {colonist.gender_subject} used to be a reporter for a large newspaper. She doesn't talk about her past though. Nobody knows why.",
          modifiers: [["memory", 4]],
        },
        {
          type: "empathy",
          description:
            "{colonist.name} feels what others feel and she can tell what people are feeling just by looking at them. This sometimes gets her into trouble, but it's one of the reasons why she's still alive.",
          modifiers: [["social", 6]],
        },
        {
          type: "superiority_complex",
          description:
            "{colonist.first_name} has a superiority complex. It doesn't help that {colonist.gender_subject} is beautiful, intelligent and takes absolutely no crap from anyone. Everyone respects {colonist.gender_object} and no one would dare try to harm {colonist.gender_object}.",
          modifier_conflicts: [
            ["intelligence", "-"],
            ["appearance", "-"],
          ],
        },
        {
          type: "facial_birthmark",
          description:
            "{colonist.first_name} has a small birthmark on {colonist.gender_object} right cheek that looks like a claw. {colonist.first_name} always wanted it removed, but couldn't afford the surgery so she learned to live with it.",
        },
        {
          type: "research_subject",
          descriptions: [
            "{colonist.first_name} doesn't know what happened to {colonist.gender_object} at birth or what was done to {colonist.gender_object} afterwards, but {colonist.gender_subject} became cold and emotionless due to being experimented on and abused.",
            "{colonist.first_name} grew up as a tortured research subject. She was experimented on to see what effect pain would have on someone over an extended period of time. She grew up without any feeling in her legs.",
          ],
        },
      ],
    },
  ];
}
```
