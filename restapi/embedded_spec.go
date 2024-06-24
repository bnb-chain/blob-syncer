// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for handling blob query in the Blob Hub.",
    "title": "Blob Hub Service API",
    "version": "1.0.0"
  },
  "host": "blob-hub",
  "paths": {
    "/": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blob"
        ],
        "summary": "Get BSC blob sidecars by block num",
        "operationId": "getBSCBlobSidecarsByBlockNum",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RPCRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/RPCResponse"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/eth/v1/beacon/blob_sidecars/{block_id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blob"
        ],
        "summary": "Get blob sidecars by block num",
        "operationId": "getBlobSidecarsByBlockNum",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "Block identifier. Can be one of: 'head' (canonical head in node's view), 'genesis', 'finalized', \u003cslot\u003e, \u003chex encoded blockRoot with 0x prefix\u003e",
            "name": "block_id",
            "in": "path",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "Array of indices for blob sidecars to request for in the specified block. Returns all blob sidecars in the block if not specified",
            "name": "indices",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/GetBlobSideCarsResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "blob not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BSCBlobSidecar": {
      "type": "object",
      "properties": {
        "blobs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "commitments": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "proofs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "BSCBlobTxSidecar": {
      "type": "object",
      "properties": {
        "blobSidecar": {
          "$ref": "#/definitions/BSCBlobSidecar"
        },
        "blockHash": {
          "type": "string"
        },
        "blockNumber": {
          "type": "string"
        },
        "txHash": {
          "type": "string"
        },
        "txIndex": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "HTTP error code",
          "type": "integer",
          "format": "int64",
          "x-omitempty": false,
          "example": "400/500"
        },
        "message": {
          "description": "Error message",
          "type": "string",
          "x-omitempty": false,
          "example": "Bad request/Internal server error"
        }
      }
    },
    "GetBlobSideCarsResponse": {
      "type": "object",
      "properties": {
        "code": {
          "description": "status code",
          "type": "integer",
          "example": 200
        },
        "data": {
          "description": "actual data for request",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Sidecar"
          }
        },
        "message": {
          "description": "error message if there is error",
          "type": "string",
          "example": "signature invalid"
        }
      }
    },
    "RPCError": {
      "type": "object",
      "properties": {
        "code": {
          "description": "RPC error code",
          "type": "integer",
          "format": "int64",
          "x-omitempty": false,
          "example": -32602
        },
        "message": {
          "description": "Error message",
          "type": "string",
          "x-omitempty": false,
          "example": "Invalid params"
        }
      }
    },
    "RPCRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "example": 1
        },
        "jsonrpc": {
          "type": "string",
          "example": "2.0"
        },
        "method": {
          "type": "string",
          "example": "eth_getBlobSidecars"
        },
        "params": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "0x1",
            true
          ]
        }
      }
    },
    "RPCResponse": {
      "type": "object",
      "properties": {
        "error": {
          "$ref": "#/definitions/RPCError"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "jsonrpc": {
          "type": "string",
          "example": "2.0"
        },
        "result": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BSCBlobTxSidecar"
          }
        }
      }
    },
    "Sidecar": {
      "type": "object",
      "properties": {
        "blob": {
          "type": "string"
        },
        "index": {
          "type": "string",
          "example": 1
        },
        "kzg_commitment": {
          "type": "string"
        },
        "kzg_commitment_inclusion_proof": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-omitempty": true
        },
        "kzg_proof": {
          "type": "string"
        },
        "signed_block_header": {
          "type": "object",
          "properties": {
            "message": {
              "type": "object",
              "properties": {
                "body_root": {
                  "type": "string"
                },
                "parent_root": {
                  "type": "string"
                },
                "proposer_index": {
                  "type": "string"
                },
                "slot": {
                  "type": "string"
                },
                "state_root": {
                  "type": "string"
                }
              }
            },
            "signature": {
              "type": "string"
            }
          }
        },
        "tx_hash": {
          "type": "string"
        },
        "tx_index": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": true
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for handling blob query in the Blob Hub.",
    "title": "Blob Hub Service API",
    "version": "1.0.0"
  },
  "host": "blob-hub",
  "paths": {
    "/": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blob"
        ],
        "summary": "Get BSC blob sidecars by block num",
        "operationId": "getBSCBlobSidecarsByBlockNum",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RPCRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/RPCResponse"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/eth/v1/beacon/blob_sidecars/{block_id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blob"
        ],
        "summary": "Get blob sidecars by block num",
        "operationId": "getBlobSidecarsByBlockNum",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "Block identifier. Can be one of: 'head' (canonical head in node's view), 'genesis', 'finalized', \u003cslot\u003e, \u003chex encoded blockRoot with 0x prefix\u003e",
            "name": "block_id",
            "in": "path",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "Array of indices for blob sidecars to request for in the specified block. Returns all blob sidecars in the block if not specified",
            "name": "indices",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/GetBlobSideCarsResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "blob not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BSCBlobSidecar": {
      "type": "object",
      "properties": {
        "blobs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "commitments": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "proofs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "BSCBlobTxSidecar": {
      "type": "object",
      "properties": {
        "blobSidecar": {
          "$ref": "#/definitions/BSCBlobSidecar"
        },
        "blockHash": {
          "type": "string"
        },
        "blockNumber": {
          "type": "string"
        },
        "txHash": {
          "type": "string"
        },
        "txIndex": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "HTTP error code",
          "type": "integer",
          "format": "int64",
          "x-omitempty": false,
          "example": "400/500"
        },
        "message": {
          "description": "Error message",
          "type": "string",
          "x-omitempty": false,
          "example": "Bad request/Internal server error"
        }
      }
    },
    "GetBlobSideCarsResponse": {
      "type": "object",
      "properties": {
        "code": {
          "description": "status code",
          "type": "integer",
          "example": 200
        },
        "data": {
          "description": "actual data for request",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Sidecar"
          }
        },
        "message": {
          "description": "error message if there is error",
          "type": "string",
          "example": "signature invalid"
        }
      }
    },
    "RPCError": {
      "type": "object",
      "properties": {
        "code": {
          "description": "RPC error code",
          "type": "integer",
          "format": "int64",
          "x-omitempty": false,
          "example": -32602
        },
        "message": {
          "description": "Error message",
          "type": "string",
          "x-omitempty": false,
          "example": "Invalid params"
        }
      }
    },
    "RPCRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "example": 1
        },
        "jsonrpc": {
          "type": "string",
          "example": "2.0"
        },
        "method": {
          "type": "string",
          "example": "eth_getBlobSidecars"
        },
        "params": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "0x1",
            true
          ]
        }
      }
    },
    "RPCResponse": {
      "type": "object",
      "properties": {
        "error": {
          "$ref": "#/definitions/RPCError"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "jsonrpc": {
          "type": "string",
          "example": "2.0"
        },
        "result": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BSCBlobTxSidecar"
          }
        }
      }
    },
    "Sidecar": {
      "type": "object",
      "properties": {
        "blob": {
          "type": "string"
        },
        "index": {
          "type": "string",
          "example": 1
        },
        "kzg_commitment": {
          "type": "string"
        },
        "kzg_commitment_inclusion_proof": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-omitempty": true
        },
        "kzg_proof": {
          "type": "string"
        },
        "signed_block_header": {
          "type": "object",
          "properties": {
            "message": {
              "type": "object",
              "properties": {
                "body_root": {
                  "type": "string"
                },
                "parent_root": {
                  "type": "string"
                },
                "proposer_index": {
                  "type": "string"
                },
                "slot": {
                  "type": "string"
                },
                "state_root": {
                  "type": "string"
                }
              }
            },
            "signature": {
              "type": "string"
            }
          }
        },
        "tx_hash": {
          "type": "string"
        },
        "tx_index": {
          "type": "integer",
          "format": "int64",
          "x-omitempty": true
        }
      }
    },
    "SidecarSignedBlockHeader": {
      "type": "object",
      "properties": {
        "message": {
          "type": "object",
          "properties": {
            "body_root": {
              "type": "string"
            },
            "parent_root": {
              "type": "string"
            },
            "proposer_index": {
              "type": "string"
            },
            "slot": {
              "type": "string"
            },
            "state_root": {
              "type": "string"
            }
          }
        },
        "signature": {
          "type": "string"
        }
      }
    },
    "SidecarSignedBlockHeaderMessage": {
      "type": "object",
      "properties": {
        "body_root": {
          "type": "string"
        },
        "parent_root": {
          "type": "string"
        },
        "proposer_index": {
          "type": "string"
        },
        "slot": {
          "type": "string"
        },
        "state_root": {
          "type": "string"
        }
      }
    }
  }
}`))
}
