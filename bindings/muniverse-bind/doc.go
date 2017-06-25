// Command muniverse-bind uses stdin and stdout to serve
// an API for controlling muniverse environments.
// It is intended to be executed by muniverse packages in
// other programming languages.
//
// The API uses a BSON protocol over stdin and stdout.
// In this protocol, the front-end sends calls to the
// back-end and receives responses back.
// Multiple calls can be running asynchronously.
//
// Call Format
//
// When a call is sent to the back-end, it is encoded as a
// BSON object with a preceding length field.
// The length field is a 4-byte little-endian integer
// indicating the size of the BSON object.
// Call objects have an "ID" field which is a string used
// to identify the response to the call.
// Call objects have a single field besides the "ID" field
// which stores the arguments of the call.
// The name of this field is the name of the call.
//
// Response Format
//
// The response to a call is, like a call, encoded as a
// BSON object with a preceding length field.
// A response object has an "ID" field matching the "ID"
// of the call.
// If a call fails, it has a non-null Error field which is
// a string error message.
// Otherwise, the response object may have other fields
// depending on the type of the return value.
//
// Types of Calls
//
// The SpecForName call looks up an environment by name.
// The call takes a "Name" string argument.
// The response has a "Spec" field which is either null or
// an environment specification object.
//
// The NewEnv call creates an environment from an
// environment specification object.
// The call takes a "Spec" argument, which is an
// environment specification object.
// The response has a "UID" field, which is a string
// identifier for the new environment instance.
//
// The NewEnvContainer call is like NewEnv, but it uses a
// custom Docker container.
// The call and response are the same as NewEnv, except
// that the call has an extra "Container" argument naming
// the Docker container.
//
// The NewEnvChrome call is like NewEnv, but it uses a
// custom Chrome instance and game server.
// The call and response are the same as NewEnv, except
// that the call has "Host" and "GameHost" arguments.
//
// The CloseEnv closes an environment.
// The call takes a "UID" argument.
// The response has no fields.
package main
