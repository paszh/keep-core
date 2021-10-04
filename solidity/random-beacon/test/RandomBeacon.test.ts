import { ethers } from "hardhat"
import { Signer, Contract } from "ethers"
import { expect } from "chai"

describe("RandomBeacon", () => {
  let governance: Signer
  let thirdParty: Signer
  let randomBeacon: Contract

  beforeEach(async () => {
    const signers = await ethers.getSigners()
    governance = signers[0]
    thirdParty = signers[1]

    const RandomBeacon = await ethers.getContractFactory("RandomBeacon")
    randomBeacon = await RandomBeacon.deploy()
    await randomBeacon.deployed()
  })

  describe("updateRelayEntryParameters", () => {
    const relayRequestFee = 100
    const relayEntrySubmissionEligibilityDelay = 200
    const RelayEntryHardTimeout = 300
    const callbackGasLimit = 400
    context("when the caller is not the owner", () => {
      it("should revert", async () => {
        await expect(
          randomBeacon
            .connect(thirdParty)
            .updateRelayEntryParameters(
              relayRequestFee,
              relayEntrySubmissionEligibilityDelay,
              RelayEntryHardTimeout,
              callbackGasLimit
            )
        ).to.be.revertedWith("Ownable: caller is not the owner")
      })
    })

    context("when the caller is the owner", () => {
      beforeEach(async () => {
        await randomBeacon
          .connect(governance)
          .updateRelayEntryParameters(
            relayRequestFee,
            relayEntrySubmissionEligibilityDelay,
            RelayEntryHardTimeout,
            callbackGasLimit
          )
      })

      it("should update the relay request fee", async () => {
        expect(await randomBeacon.relayRequestFee()).to.be.equal(
          relayRequestFee
        )
      })

      it("should update the relay entry submission eligibility delay", async () => {
        expect(
          await randomBeacon.relayEntrySubmissionEligibilityDelay()
        ).to.be.equal(relayEntrySubmissionEligibilityDelay)
      })

      it("should update the relay entry hard timeout", async () => {
        expect(await randomBeacon.relayEntryHardTimeout()).to.be.equal(
          RelayEntryHardTimeout
        )
      })

      it("should update the callback gas limit", async () => {
        expect(await randomBeacon.callbackGasLimit()).to.be.equal(
          callbackGasLimit
        )
      })
    })
  })

  describe("updateGroupCreationParameters", () => {
    const groupCreationFrequency = 100
    const groupLifetime = 200
    const dkgResultChallengePeriodLength = 300
    const dkgSubmissionEligibilityDelay = 400
    context("when the caller is not the owner", () => {
      it("should revert", async () => {
        await expect(
          randomBeacon
            .connect(thirdParty)
            .updateGroupCreationParameters(
              groupCreationFrequency,
              groupLifetime,
              dkgResultChallengePeriodLength,
              dkgSubmissionEligibilityDelay
            )
        ).to.be.revertedWith("Ownable: caller is not the owner")
      })
    })

    context("when the caller is the owner", () => {
      beforeEach(async () => {
        await randomBeacon
          .connect(governance)
          .updateGroupCreationParameters(
            groupCreationFrequency,
            groupLifetime,
            dkgResultChallengePeriodLength,
            dkgSubmissionEligibilityDelay
          )
      })

      it("should update the group creation frequency", async () => {
        expect(await randomBeacon.groupCreationFrequency()).to.be.equal(
          groupCreationFrequency
        )
      })

      it("should update the group lifetime", async () => {
        expect(await randomBeacon.groupLifetime()).to.be.equal(groupLifetime)
      })

      it("should update the DKG result challenge period length", async () => {
        expect(await randomBeacon.dkgResultChallengePeriodLength()).to.be.equal(
          dkgResultChallengePeriodLength
        )
      })

      it("should update the DKG submission eligibility delay", async () => {
        expect(await randomBeacon.dkgSubmissionEligibilityDelay()).to.be.equal(
          dkgSubmissionEligibilityDelay
        )
      })
    })
  })

  describe("updateRewardParameters", () => {
    const dkgResultSubmissionReward = 100
    const sortitionPoolUnlockingReward = 200
    context("when the caller is not the owner", () => {
      it("should revert", async () => {
        await expect(
          randomBeacon
            .connect(thirdParty)
            .updateRewardParameters(
              dkgResultSubmissionReward,
              sortitionPoolUnlockingReward
            )
        ).to.be.revertedWith("Ownable: caller is not the owner")
      })
    })

    context("when the caller is the owner", () => {
      beforeEach(async () => {
        await randomBeacon
          .connect(governance)
          .updateRewardParameters(
            dkgResultSubmissionReward,
            sortitionPoolUnlockingReward
          )
      })

      it("should update the DKG result submission reward", async () => {
        expect(await randomBeacon.dkgResultSubmissionReward()).to.be.equal(
          dkgResultSubmissionReward
        )
      })

      it("should update the sortition pool unlocking reward", async () => {
        expect(await randomBeacon.sortitionPoolUnlockingReward()).to.be.equal(
          sortitionPoolUnlockingReward
        )
      })
    })
  })

  describe("updateSlashingParameters", () => {
    const relayEntrySubmissionFailureSlashingAmount = 100
    const maliciousDkgResultSlashingAmount = 200
    context("when the caller is not the owner", () => {
      it("should revert", async () => {
        await expect(
          randomBeacon
            .connect(thirdParty)
            .updateSlashingParameters(
              relayEntrySubmissionFailureSlashingAmount,
              maliciousDkgResultSlashingAmount
            )
        ).to.be.revertedWith("Ownable: caller is not the owner")
      })
    })

    context("when the caller is the owner", () => {
      beforeEach(async () => {
        await randomBeacon
          .connect(governance)
          .updateSlashingParameters(
            relayEntrySubmissionFailureSlashingAmount,
            maliciousDkgResultSlashingAmount
          )
      })

      it("should update the relay entry submission failure slashing amount", async () => {
        expect(
          await randomBeacon.relayEntrySubmissionFailureSlashingAmount()
        ).to.be.equal(relayEntrySubmissionFailureSlashingAmount)
      })

      it("should update the malicious DKG result slashing amount", async () => {
        expect(
          await randomBeacon.maliciousDkgResultSlashingAmount()
        ).to.be.equal(maliciousDkgResultSlashingAmount)
      })
    })
  })
})
