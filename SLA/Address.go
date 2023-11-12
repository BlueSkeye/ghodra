package ghodra;


/// \brief A low-level machine address for labelling bytes and data.
///
/// All data that can be manipulated within the processor reverse
/// engineering model can be labelled with an Address. It is
/// simply an address space (AddrSpace) and an offset within that
/// space.  Note that processor registers are typically modelled
/// by creating a dedicated address space for them, as distinct
/// from RAM say, and then specifying certain addresses within the
/// register space that correspond to particular registers. However,
/// an arbitrary address could refer to anything,
/// RAM, ROM, cpu register, data segment, coprocessor, stack,
/// nvram, etc.
/// An Address represents an offset \e only, not an offset and length
type Address struct {
// protected:
	AddrSpace *base; ///< Pointer to our address space
	uintb offset; ///< Offset (in bytes)
}

///< Initialize an extremal address
/// Some data structures sort on an Address, and it is convenient
/// to be able to create an Address that is either bigger than
/// or smaller than all other Addresses.
/// \param ex is either \e m_minimal or \e m_maximal
func NewAddress(mach_extreme ex) Address {
	var result *Address
	result := &Address()
	if ex == m_minimal {
	  base = nil;
	  offset = 0;
	} else {
	  base = (AddrSpace*) ~((uintp)0);
	  offset = ~((uintb)0);
	}
	return result;
}

//   Address(void);		///< Create an invalid address
//   Address(AddrSpace *id,uintb off); ///< Construct an address with a space/offset pair
//   Address(const Address &op2);	///< A copy constructor

//   bool isInvalid(void) const;  ///< Is the address invalid?
//   int4 getAddrSize(void) const; ///< Get the number of bytes in the address
//   bool isBigEndian(void) const;	///< Is data at this address big endian encoded
//   void printRaw(ostream &s) const; ///< Write a raw version of the address to a stream
//   int4 read(const string &s); ///< Read in the address from a string
//   AddrSpace *getSpace(void) const; ///< Get the address space
//   uintb getOffset(void) const;  ///< Get the address offset
//   char getShortcut(void) const;	///< Get the shortcut character for the address space
//   Address &operator=(const Address &op2); ///< Copy an address
//   bool operator==(const Address &op2) const; ///< Compare two addresses for equality
//   bool operator!=(const Address &op2) const; ///< Compare two addresses for inequality
//   bool operator<(const Address &op2) const; ///< Compare two addresses via their natural ordering
//   bool operator<=(const Address &op2) const; ///< Compare two addresses via their natural ordering
//   Address operator+(int8 off) const; ///< Increment address by a number of bytes
//   Address operator-(int8 off) const; ///< Decrement address by a number of bytes
//   friend ostream &operator<<(ostream &s,const Address &addr);  ///< Write out an address to stream
//   bool containedBy(int4 sz,const Address &op2,int4 sz2) const;	///< Determine if \e op2 range contains \b this range
//   int4 justifiedContain(int4 sz,const Address &op2,int4 sz2,bool forceleft) const; ///< Determine if \e op2 is the least significant part of \e this.
//   int4 overlap(int4 skip,const Address &op,int4 size) const; ///< Determine how \b this address falls in a given address range
//   int4 overlapJoin(int4 skip,const Address &op,int4 size) const;	///< Determine how \b this falls in a possible \e join space address range
//   bool isContiguous(int4 sz,const Address &loaddr,int4 losz) const; ///< Does \e this form a contiguous range with \e loaddr
//   bool isConstant(void) const; ///< Is this a \e constant \e value
//   void renormalize(int4 size);	///< Make sure there is a backing JoinRecord if \b this is in the \e join space
//   bool isJoin(void) const;	///< Is this a \e join \e value
//   void encode(Encoder &encoder) const; ///< Encode \b this to a stream
//   void encode(Encoder &encoder,int4 size) const; ///< Encode \b this and a size to a stream

//   /// Restore an address from parsed XML
//   static Address decode(Decoder &decoder);

//   /// Restore an address and size from parsed XML
//   static Address decode(Decoder &decoder,int4 &size);
