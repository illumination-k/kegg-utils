package kegg

import (
	"kegg/hashset"
	"kegg/kgml"
)

type Orthology struct {
	ID          string
	Name        string
	ReactionIDs hashset.HashSet[string]
	PathwayIDs  hashset.HashSet[string]
}

type Pathway struct {
	ID           string
	Name         string
	PathwayIDs   hashset.HashSet[string]
	ReactionIDs  hashset.HashSet[string]
	OrthologyIDs hashset.HashSet[string]
}

type Reaction struct {
	ID           string
	Name         string
	PathwayIDs   hashset.HashSet[string]
	OrthologyIDs hashset.HashSet[string]
}

type KeggValueContainer struct {
	orthologyIdToName map[string]string
	reactionIdToName  map[string]string
	PathwayMap        map[string]Pathway
	OrthologyMap      map[string]Orthology
	ReactionMap       map[string]Reaction
}

func NewKeggValueContainer(
	orthologyIdToName map[string]string,
	reactionIdToName map[string]string,
) KeggValueContainer {
	return KeggValueContainer{
		orthologyIdToName: orthologyIdToName,
		reactionIdToName:  reactionIdToName,
	}
}

func (c *KeggValueContainer) AddKGML(
	kgml kgml.KGMLPathway,
) {
	pathwayId := kgml.Name
	pathway, found := c.PathwayMap[pathwayId]

	if !found {
		pathway = Pathway{ID: pathwayId, Name: kgml.Title}
	}

	for _, entry := range kgml.Entries {
		if entry.Type == "pathway" {
			pathway.PathwayIDs.Add(entry.Names...)
		} else if entry.Type == "orthology" {
			// Add ReactionID and Pathway ID into Pathay
			pathway.OrthologyIDs.Add(entry.Names...)
			pathway.ReactionIDs.Add(entry.Reaction)

			// Set reaction
			reactionId := entry.Reaction

			reaction, found := c.ReactionMap[reactionId]
			if !found {
				reaction = Reaction{ID: reactionId, Name: c.reactionIdToName[reactionId]}
			}

			reaction.OrthologyIDs.Add(entry.Names...)
			reaction.PathwayIDs.Add(pathway.ID)
			c.ReactionMap[reactionId] = reaction

			for _, orthologyId := range entry.Names {
				orthology, found := c.OrthologyMap[orthologyId]

				if !found {
					orthology = Orthology{ID: orthologyId, Name: c.orthologyIdToName[orthologyId]}
				}

				orthology.PathwayIDs.Add(pathway.ID)
				orthology.ReactionIDs.Add(reactionId)
				c.OrthologyMap[orthologyId] = orthology
			}
		}
	}

	c.PathwayMap[pathwayId] = pathway
}
