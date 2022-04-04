package main

const RuleSetSchema = `{
    "$schema": "https://json-schema.org/draft-07/schema#",
    "title": "Linter Rules",
    "description": "Rules for validating Genesys Cloud Developer Center content.",
    "type": "object",
    "properties": {
        "name": {
            "description": "Name of the rule configuration",
            "type": "string"
        },
        "description": {
            "description": "Description of the rule configuration",
            "type": "string"
        },
        "ruleGroups": {
            "description": "Groups of rules based on common aspect of validation.",
            "type": "object",
            "patternProperties": {
                "^[A-Z]+$":{
                    "description": "ID of the rule group. This will also be the prefix for the specific rules's id.",
                    "type": "object",
                    "properties": {
                        "description": {
                            "description": "Description of the rule group.",
                            "type": "string"
                        },
                        "rules": {
                            "description": "Rules for validation",
                            "type": "object",
                            "patternProperties": {
                                "[0-9]+": {
                                    "description": "Defines the rule to validate",
                                    "type": "object",
                                    "properties": {
                                        "description": {
                                            "description": "Description of the rule.",
                                            "type": "string"
                                        },
                                        "path": {
                                            "description": "Path for the file/folder to be evaluated.",
                                            "type": "string",
                                            "pattern": "^(.+)/([^/]+)$"
                                        },
                                        "files": {
                                            "description": "Array of files/folder to be evaluated.",
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        },
                                        "conditions": {
                                            "description": "Array of conditions to evaluate against. All conditions must pass for the rule to pass.",
                                            "type": "array",
                                            "items": {
                                                "description": "A condition to evaluate on the rule.",
                                                "type": "object",
                                                "properties": {
                                                    "pathExists": {
                                                        "description": "Check whether the path(file/folder) exists. Setting this to false does not have an effect.",
                                                        "enum": [true]
                                                    },
                                                    "contains": {
                                                        "description": "Checks the plaintext file if it contains a specific value.",
                                                        "type": "array",
                                                        "items": {
                                                            "description": "Definition for content to find",
                                                            "type": "object",
                                                            "properties": {
                                                                "type": {
                                                                    "description": "Valid: static, regex",
                                                                    "enum": ["static", "regex"]
                                                                },
                                                                "value": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": ["type", "value"]
                                                        }
                                                    },
                                                    "notContains": {
                                                        "description": "Checks the plaintext file that nothing matches the regex pattern.",
                                                        "type": "array",
                                                        "items": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "checkReferenceExist": {
                                                        "description": "Checks if the path(file/folder) exists in the blueprints",
                                                        "type": "array",
                                                        "items": {
                                                            "type": "string"
                                                        }
                                                    }
                                                },
                                                "additionalProperties": false
                                            },
                                            "minItems": 1
                                        },
                                        "level": {
                                            "description": "Severity level of the rule. Valid: ['warning', 'error']",
                                            "enum": ["warning", "error"]
                                        }
                                    },
                                    "required": ["description", "conditions", "level"],
                                    "oneOf": [{
                                        "required": ["files"],
                                        "properties": {
                                            "path": false 
                                        }
                                    }, {
                                        "required": ["path"],
                                        "properties": {
                                            "files": false 
                                        }
                                    }],
                                    "additionalProperties": false
                                }
                            },
                            "additionalProperties": false
                        }
                    },
                    "required": ["rules"],
                    "additionalProperties": false
                }
            },
            "additionalProperties": false
        }
    },
    "required": ["name", "description", "ruleGroups"],
    "additionalProperties": false
}

`
