package ghodra

/// An enum for specifying extremal addresses
type mach_extreme int64

const (
	m_minimal mach_extreme = iota ///< Smallest possible address
	m_maximal                     ///< Biggest possible address
)
