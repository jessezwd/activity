package propertypublickey

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// W3IDSecurityV1PublicKeyPropertyIterator is an iterator for a property. It is
// permitted to be a single nilable value type.
type W3IDSecurityV1PublicKeyPropertyIterator struct {
	w3idsecurityv1PublicKeyMember vocab.W3IDSecurityV1PublicKey
	unknown                       interface{}
	iri                           *url.URL
	alias                         string
	myIdx                         int
	parent                        vocab.W3IDSecurityV1PublicKeyProperty
}

// NewW3IDSecurityV1PublicKeyPropertyIterator creates a new
// W3IDSecurityV1PublicKey property.
func NewW3IDSecurityV1PublicKeyPropertyIterator() *W3IDSecurityV1PublicKeyPropertyIterator {
	return &W3IDSecurityV1PublicKeyPropertyIterator{alias: "widsv"}
}

// deserializeW3IDSecurityV1PublicKeyPropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeW3IDSecurityV1PublicKeyPropertyIterator(i interface{}, aliasMap map[string]string) (*W3IDSecurityV1PublicKeyPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://w3id.org/security/v1"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &W3IDSecurityV1PublicKeyPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializePublicKeyW3IDSecurityV1()(m, aliasMap); err == nil {
			this := &W3IDSecurityV1PublicKeyPropertyIterator{
				alias:                         alias,
				w3idsecurityv1PublicKeyMember: v,
			}
			return this, nil
		}
	}
	this := &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// Get returns the value of this property. When IsW3IDSecurityV1PublicKey returns
// false, Get will return any arbitrary value.
func (this W3IDSecurityV1PublicKeyPropertyIterator) Get() vocab.W3IDSecurityV1PublicKey {
	return this.w3idsecurityv1PublicKeyMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this W3IDSecurityV1PublicKeyPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetType returns the value in this property as a Type. Returns nil if the value
// is not an ActivityStreams type, such as an IRI or another value.
func (this W3IDSecurityV1PublicKeyPropertyIterator) GetType() vocab.Type {
	if this.IsW3IDSecurityV1PublicKey() {
		return this.Get()
	}

	return nil
}

// HasAny returns true if the value or IRI is set.
func (this W3IDSecurityV1PublicKeyPropertyIterator) HasAny() bool {
	return this.IsW3IDSecurityV1PublicKey() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this W3IDSecurityV1PublicKeyPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsW3IDSecurityV1PublicKey returns true if this property is set and not an IRI.
func (this W3IDSecurityV1PublicKeyPropertyIterator) IsW3IDSecurityV1PublicKey() bool {
	return this.w3idsecurityv1PublicKeyMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this W3IDSecurityV1PublicKeyPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://w3id.org/security/v1": this.alias}
	var child map[string]string
	if this.IsW3IDSecurityV1PublicKey() {
		child = this.Get().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this W3IDSecurityV1PublicKeyPropertyIterator) KindIndex() int {
	if this.IsW3IDSecurityV1PublicKey() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this W3IDSecurityV1PublicKeyPropertyIterator) LessThan(o vocab.W3IDSecurityV1PublicKeyPropertyIterator) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsW3IDSecurityV1PublicKey() && !o.IsW3IDSecurityV1PublicKey() {
		// Both are unknowns.
		return false
	} else if this.IsW3IDSecurityV1PublicKey() && !o.IsW3IDSecurityV1PublicKey() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsW3IDSecurityV1PublicKey() && o.IsW3IDSecurityV1PublicKey() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return this.Get().LessThan(o.Get())
	}
}

// Name returns the name of this property: "W3IDSecurityV1PublicKey".
func (this W3IDSecurityV1PublicKeyPropertyIterator) Name() string {
	return "W3IDSecurityV1PublicKey"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this W3IDSecurityV1PublicKeyPropertyIterator) Next() vocab.W3IDSecurityV1PublicKeyPropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this W3IDSecurityV1PublicKeyPropertyIterator) Prev() vocab.W3IDSecurityV1PublicKeyPropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// Set sets the value of this property. Calling IsW3IDSecurityV1PublicKey
// afterwards will return true.
func (this *W3IDSecurityV1PublicKeyPropertyIterator) Set(v vocab.W3IDSecurityV1PublicKey) {
	this.clear()
	this.w3idsecurityv1PublicKeyMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *W3IDSecurityV1PublicKeyPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetType attempts to set the property for the arbitrary type. Returns an error
// if it is not a valid type to set on this property.
func (this *W3IDSecurityV1PublicKeyPropertyIterator) SetType(t vocab.Type) error {
	if v, ok := t.(vocab.W3IDSecurityV1PublicKey); ok {
		this.Set(v)
		return nil
	}

	return fmt.Errorf("illegal type to set on W3IDSecurityV1PublicKey property: %T", t)
}

// clear ensures no value of this property is set. Calling
// IsW3IDSecurityV1PublicKey afterwards will return false.
func (this *W3IDSecurityV1PublicKeyPropertyIterator) clear() {
	this.unknown = nil
	this.iri = nil
	this.w3idsecurityv1PublicKeyMember = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this W3IDSecurityV1PublicKeyPropertyIterator) serialize() (interface{}, error) {
	if this.IsW3IDSecurityV1PublicKey() {
		return this.Get().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// W3IDSecurityV1PublicKeyProperty is the non-functional property "publicKey". It
// is permitted to have one or more values, and of different value types.
type W3IDSecurityV1PublicKeyProperty struct {
	properties []*W3IDSecurityV1PublicKeyPropertyIterator
	alias      string
}

// DeserializePublicKeyProperty creates a "publicKey" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializePublicKeyProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.W3IDSecurityV1PublicKeyProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://w3id.org/security/v1"]; ok {
		alias = a
	}
	propName := "publicKey"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "publicKey")
	}
	i, ok := m[propName]

	if ok {
		this := &W3IDSecurityV1PublicKeyProperty{
			alias:      alias,
			properties: []*W3IDSecurityV1PublicKeyPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeW3IDSecurityV1PublicKeyPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeW3IDSecurityV1PublicKeyPropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewW3IDSecurityV1PublicKeyProperty creates a new publicKey property.
func NewW3IDSecurityV1PublicKeyProperty() *W3IDSecurityV1PublicKeyProperty {
	return &W3IDSecurityV1PublicKeyProperty{alias: "widsv"}
}

// AppendIRI appends an IRI value to the back of a list of the property "publicKey"
func (this *W3IDSecurityV1PublicKeyProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "publicKey". Invalidates iterators that are traversing using Prev.
// Returns an error if the type is not a valid one to set for this property.
func (this *W3IDSecurityV1PublicKeyProperty) AppendType(t vocab.Type) error {
	n := &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:  this.alias,
		myIdx:  this.Len(),
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append(this.properties, n)
	return nil
}

// AppendW3IDSecurityV1PublicKey appends a PublicKey value to the back of a list
// of the property "publicKey". Invalidates iterators that are traversing
// using Prev.
func (this *W3IDSecurityV1PublicKeyProperty) AppendW3IDSecurityV1PublicKey(v vocab.W3IDSecurityV1PublicKey) {
	this.properties = append(this.properties, &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:                         this.alias,
		myIdx:                         this.Len(),
		parent:                        this,
		w3idsecurityv1PublicKeyMember: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this W3IDSecurityV1PublicKeyProperty) At(index int) vocab.W3IDSecurityV1PublicKeyPropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this W3IDSecurityV1PublicKeyProperty) Begin() vocab.W3IDSecurityV1PublicKeyPropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this W3IDSecurityV1PublicKeyProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this W3IDSecurityV1PublicKeyProperty) End() vocab.W3IDSecurityV1PublicKeyPropertyIterator {
	return nil
}

// Insert inserts an IRI value at the specified index for a property "publicKey".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *W3IDSecurityV1PublicKeyProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "publicKey". Invalidates all iterators. Returns an error if the
// type is not a valid one to set for this property.
func (this *W3IDSecurityV1PublicKeyProperty) InsertType(idx int, t vocab.Type) error {
	n := &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = n
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
	return nil
}

// InsertW3IDSecurityV1PublicKey inserts a PublicKey value at the specified index
// for a property "publicKey". Existing elements at that index and higher are
// shifted back once. Invalidates all iterators.
func (this *W3IDSecurityV1PublicKeyProperty) InsertW3IDSecurityV1PublicKey(idx int, v vocab.W3IDSecurityV1PublicKey) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:                         this.alias,
		myIdx:                         idx,
		parent:                        this,
		w3idsecurityv1PublicKeyMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this W3IDSecurityV1PublicKeyProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://w3id.org/security/v1": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this W3IDSecurityV1PublicKeyProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "publicKey" property.
func (this W3IDSecurityV1PublicKeyProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this W3IDSecurityV1PublicKeyProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].Get()
			rhs := this.properties[j].Get()
			return lhs.LessThan(rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this W3IDSecurityV1PublicKeyProperty) LessThan(o vocab.W3IDSecurityV1PublicKeyProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property: "publicKey".
func (this W3IDSecurityV1PublicKeyProperty) Name() string {
	return "publicKey"
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "publicKey".
func (this *W3IDSecurityV1PublicKeyProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*W3IDSecurityV1PublicKeyPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "publicKey". Invalidates all iterators. Returns an error if the
// type is not a valid one to set for this property.
func (this *W3IDSecurityV1PublicKeyProperty) PrependType(t vocab.Type) error {
	n := &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:  this.alias,
		myIdx:  0,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append([]*W3IDSecurityV1PublicKeyPropertyIterator{n}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
	return nil
}

// PrependW3IDSecurityV1PublicKey prepends a PublicKey value to the front of a
// list of the property "publicKey". Invalidates all iterators.
func (this *W3IDSecurityV1PublicKeyProperty) PrependW3IDSecurityV1PublicKey(v vocab.W3IDSecurityV1PublicKey) {
	this.properties = append([]*W3IDSecurityV1PublicKeyPropertyIterator{{
		alias:                         this.alias,
		myIdx:                         0,
		parent:                        this,
		w3idsecurityv1PublicKeyMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "publicKey", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *W3IDSecurityV1PublicKeyProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &W3IDSecurityV1PublicKeyPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this W3IDSecurityV1PublicKeyProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// Set sets a PublicKey value to be at the specified index for the property
// "publicKey". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *W3IDSecurityV1PublicKeyProperty) Set(idx int, v vocab.W3IDSecurityV1PublicKey) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:                         this.alias,
		myIdx:                         idx,
		parent:                        this,
		w3idsecurityv1PublicKeyMember: v,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property
// "publicKey". Panics if the index is out of bounds.
func (this *W3IDSecurityV1PublicKeyProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetType sets an arbitrary type value to the specified index of the property
// "publicKey". Invalidates all iterators. Returns an error if the type is not
// a valid one to set for this property. Panics if the index is out of bounds.
func (this *W3IDSecurityV1PublicKeyProperty) SetType(idx int, t vocab.Type) error {
	n := &W3IDSecurityV1PublicKeyPropertyIterator{
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	(this.properties)[idx] = n
	return nil
}

// Swap swaps the location of values at two indices for the "publicKey" property.
func (this W3IDSecurityV1PublicKeyProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
