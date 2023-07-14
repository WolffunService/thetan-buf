// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/shared/v1/game_info.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Shared.V1 {

  /// <summary>Holder for reflection information generated from thetan/shared/v1/game_info.proto</summary>
  public static partial class GameInfoReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/shared/v1/game_info.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static GameInfoReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiB0aGV0YW4vc2hhcmVkL3YxL2dhbWVfaW5mby5wcm90bxIQdGhldGFuLnNo",
            "YXJlZC52MSpSCghHYW1lTW9kZRIKCgZSQU5LRUQQABIRCg1TUEVDSUFMX0VW",
            "RU5UEAESDwoLQ1VTVE9NX01PREUQAhIWChJUSEVUQU5fUklWQUxTX01PREUQ",
            "Ayq1AgoKSW5HYW1lTW9kZRIVChFURUFNX0NPTExFQ1RfU1RBUhAAEhEKDVNP",
            "TE9fU1VSVklWQUwQARIRCg1EVUFMX1NVUlZJVkFMEAISHAoYVEVBTV9DT0xM",
            "RUNUX1NUQVJfNF9WU180EAMSCAoES0lORxAFEg8KC0RFQVRIX01BVENIEAYS",
            "FgoSREVBVEhfTUFUQ0hfM19WU18zEAcSCAoERkxBRxAIEgkKBVRPV0VSEAkS",
            "EQoNQkFUVExFX1JPWUFMRRAMEhcKE1NRVUFEX0JBVFRMRV9ST1lBTEUQDRIV",
            "ChFEVU9fQkFUVExFX1JPWUFMRRAOEhYKElRSSU9fQkFUVExFX1JPQVlMRRAP",
            "EhEKDVRIRVRBTl9SSVZBTFMQFBIWCglOT05FX01PREUQ////////////AUK8",
            "AQoUY29tLnRoZXRhbi5zaGFyZWQudjFCDUdhbWVJbmZvUHJvdG9QAVozdGhl",
            "dGFuLWJ1Zi9nZW4vZ28vdGhldGFuL3NoYXJlZC92MTt0aGV0YW5fc2hhcmVk",
            "X3YxogIDVFNYqgIQVGhldGFuLlNoYXJlZC5WMcoCEFRoZXRhblxTaGFyZWRc",
            "VjHiAhxUaGV0YW5cU2hhcmVkXFYxXEdQQk1ldGFkYXRh6gISVGhldGFuOjpT",
            "aGFyZWQ6OlYxYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Thetan.Shared.V1.GameMode), typeof(global::Thetan.Shared.V1.InGameMode), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum GameMode {
    [pbr::OriginalName("RANKED")] Ranked = 0,
    [pbr::OriginalName("SPECIAL_EVENT")] SpecialEvent = 1,
    [pbr::OriginalName("CUSTOM_MODE")] CustomMode = 2,
    [pbr::OriginalName("THETAN_RIVALS_MODE")] ThetanRivalsMode = 3,
  }

  public enum InGameMode {
    [pbr::OriginalName("TEAM_COLLECT_STAR")] TeamCollectStar = 0,
    [pbr::OriginalName("SOLO_SURVIVAL")] SoloSurvival = 1,
    [pbr::OriginalName("DUAL_SURVIVAL")] DualSurvival = 2,
    [pbr::OriginalName("TEAM_COLLECT_STAR_4_VS_4")] TeamCollectStar4Vs4 = 3,
    [pbr::OriginalName("KING")] King = 5,
    [pbr::OriginalName("DEATH_MATCH")] DeathMatch = 6,
    [pbr::OriginalName("DEATH_MATCH_3_VS_3")] DeathMatch3Vs3 = 7,
    [pbr::OriginalName("FLAG")] Flag = 8,
    [pbr::OriginalName("TOWER")] Tower = 9,
    [pbr::OriginalName("BATTLE_ROYALE")] BattleRoyale = 12,
    [pbr::OriginalName("SQUAD_BATTLE_ROYALE")] SquadBattleRoyale = 13,
    [pbr::OriginalName("DUO_BATTLE_ROYALE")] DuoBattleRoyale = 14,
    [pbr::OriginalName("TRIO_BATTLE_ROAYLE")] TrioBattleRoayle = 15,
    [pbr::OriginalName("THETAN_RIVALS")] ThetanRivals = 20,
    [pbr::OriginalName("NONE_MODE")] NoneMode = -1,
  }

  #endregion

}

#endregion Designer generated code