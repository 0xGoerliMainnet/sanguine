// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MultiCallable} from "./MultiCallable.sol";
import {Versioned} from "./Version.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @notice Base contract for all messaging contracts.
 * - Provides context on the local chain's domain.
 * - Provides ownership functionality.
 * - Will be providing pausing functionality when it is implemented.
 */
abstract contract MessagingBase is MultiCallable, Versioned, OwnableUpgradeable {
    // ════════════════════════════════════════════════ IMMUTABLES ═════════════════════════════════════════════════════

    /// @notice Domain of the local chain, set once upon contract creation
    uint32 public immutable localDomain;
    // @notice Domain of the Synapse chain, set once upon contract creation
    uint32 public immutable synapseDomain;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev gap for upgrade safety
    uint256[50] private __GAP; // solhint-disable-line var-name-mixedcase

    constructor(string memory version_, uint32 synapseDomain_) Versioned(version_) {
        // TODO: do we want to/need to check for overflow?
        localDomain = uint32(block.chainid);
        synapseDomain = synapseDomain_;
    }

    // TODO: Implement pausing

    /**
     * @dev Should be impossible to renounce ownership;
     * we override OpenZeppelin OwnableUpgradeable's
     * implementation of renounceOwnership to make it a no-op
     */
    function renounceOwnership() public override onlyOwner {} //solhint-disable-line no-empty-blocks
}
