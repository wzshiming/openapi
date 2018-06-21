package spec

//Expressions Runtime expressions allow defining values based on information that will only be available within the HTTP message in an actual API call.
// This mechanism is used by Link Objects and Callback Objects.
//The runtime expression is defined by the following ABNF syntax
//      expression = ( "$url" | "$method" | "$statusCode" | "$request." source | "$response." source )
//      source = ( header-reference | query-reference | path-reference | body-reference )
//      header-reference = "header." token
//      query-reference = "query." name
//      path-reference = "path." name
//      body-reference = "body" ["#" fragment]
//      fragment = a JSON Pointer [RFC 6901](https://tools.ietf.org/html/rfc6901)
//      name = *( char )
//      char = as per RFC [7159](https://tools.ietf.org/html/rfc7159#section-7)
//      token = as per RFC [7230](https://tools.ietf.org/html/rfc7230#section-3.2.6)
//The name identifier is case-sensitive, whereas token is not.
//The table below provides examples of runtime expressions and examples of their use in a value:
type Expressions string
