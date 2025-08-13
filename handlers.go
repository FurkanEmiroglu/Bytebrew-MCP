package main

import (
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateNewGameRegistry(input map[string]any, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if _, ok := input["gameId"].(string); !ok {
		return mcp.NewToolResultError("gameId is required"), nil
	}

	if _, ok := input["name"].(string); !ok {
		return mcp.NewToolResultError("name is required"), nil
	}

	if _, ok := input["gameSecretKey"].(string); !ok {
		return mcp.NewToolResultError("gameSecretKey is required"), nil
	}

	registry := GameRegistry{
		GameId:        input["gameId"].(string),
		Name:          input["name"].(string),
		GameSecretKey: input["gameSecretKey"].(string),
	}

	err := globalGameRepository.Add(registry)

	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText("Game registry created successfully"), nil
}

func ListGameRegistries(input map[string]any, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entities, err := globalGameRepository.ListEntities()

	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if len(entities) == 0 {
		return mcp.NewToolResultText("No game registries found"), nil
	}

	result := make([]map[string]any, len(entities))
	for i, entity := range entities {
		result[i] = map[string]any{
			"gameId":        entity.GameId,
			"name":          entity.Name,
			"gameSecretKey": entity.GameSecretKey,
		}
	}
	return mcp.NewToolResultStructured(result, "Structured output failed"), nil
}

func GetGameRegistry(input map[string]any, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if _, ok := input["gameId"].(string); !ok {
		return mcp.NewToolResultError("gameId is required"), nil
	}

	registry, err := globalGameRepository.Get(input["gameId"].(string))

	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultStructured(registry, "Structured output failed"), nil
}

func UpdateGameRegistry(input map[string]any, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if _, ok := input["gameId"].(string); !ok {
		return mcp.NewToolResultError("gameId is required"), nil
	}

	registry, err := globalGameRepository.Get(input["gameId"].(string))
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if name, ok := input["name"].(string); ok {
		registry.Name = name
	}

	if gameSecretKey, ok := input["gameSecretKey"].(string); ok {
		registry.GameSecretKey = gameSecretKey
	}

	err = globalGameRepository.Update(registry)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText("Game registry updated successfully"), nil
}

func DeleteGameRegistry(input map[string]any, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if _, ok := input["gameId"].(string); !ok {
		return mcp.NewToolResultError("gameId is required"), nil
	}

	err := globalGameRepository.Delete(input["gameId"].(string))
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText("Game registry deleted successfully"), nil
}
