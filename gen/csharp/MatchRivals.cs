// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/match/v1/match_rivals.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Match.V1 {

  /// <summary>Holder for reflection information generated from thetan/match/v1/match_rivals.proto</summary>
  public static partial class MatchRivalsReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/match/v1/match_rivals.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MatchRivalsReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiJ0aGV0YW4vbWF0Y2gvdjEvbWF0Y2hfcml2YWxzLnByb3RvEg90aGV0YW4u",
            "bWF0Y2gudjEaGWdvb2dsZS9wcm90b2J1Zi9hbnkucHJvdG8iVAoKTWF0Y2hS",
            "b3VuZBImCg5wbGF5ZXJJblJvdW5kcxgBIAMoCVIOcGxheWVySW5Sb3VuZHMS",
            "HgoKZ2FtZUlucHV0cxgCIAEoDFIKZ2FtZUlucHV0cyL8AQoLTWF0Y2hQbGF5",
            "ZXISGgoIcGxheWVySUQYASABKAlSCHBsYXllcklEEh4KCmJhdHRsZVJhbmsY",
            "AiABKAVSCmJhdHRsZVJhbmsSHgoKcm91bmRUaW1lcxgDIAMoAlIKcm91bmRU",
            "aW1lcxJACgZleHRyYXMYBCADKAsyKC50aGV0YW4ubWF0Y2gudjEuTWF0Y2hQ",
            "bGF5ZXIuRXh0cmFzRW50cnlSBmV4dHJhcxpPCgtFeHRyYXNFbnRyeRIQCgNr",
            "ZXkYASABKAlSA2tleRIqCgV2YWx1ZRgCIAEoCzIULmdvb2dsZS5wcm90b2J1",
            "Zi5BbnlSBXZhbHVlOgI4ASKxAQoSU3BlY3RhdG9yQmF0dGxlRW5kEhgKB21h",
            "dGNoSUQYASABKAlSB21hdGNoSUQSMwoGcm91bmRzGAIgAygLMhsudGhldGFu",
            "Lm1hdGNoLnYxLk1hdGNoUm91bmRSBnJvdW5kcxI2CgdwbGF5ZXJzGAMgAygL",
            "MhwudGhldGFuLm1hdGNoLnYxLk1hdGNoUGxheWVyUgdwbGF5ZXJzEhQKBWVu",
            "ZEF0GAQgASgDUgVlbmRBdEK6AQoTY29tLnRoZXRhbi5tYXRjaC52MUIQTWF0",
            "Y2hSaXZhbHNQcm90b1ABWjF0aGV0YW4tYnVmL2dlbi9nby90aGV0YW4vbWF0",
            "Y2gvdjE7dGhldGFuX21hdGNoX3YxogIDVE1YqgIPVGhldGFuLk1hdGNoLlYx",
            "ygIQVGhldGFuXE1hdGNoX1xWMeICHFRoZXRhblxNYXRjaF9cVjFcR1BCTWV0",
            "YWRhdGHqAhFUaGV0YW46Ok1hdGNoOjpWMWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Protobuf.WellKnownTypes.AnyReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Match.V1.MatchRound), global::Thetan.Match.V1.MatchRound.Parser, new[]{ "PlayerInRounds", "GameInputs" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Match.V1.MatchPlayer), global::Thetan.Match.V1.MatchPlayer.Parser, new[]{ "PlayerID", "BattleRank", "RoundTimes", "Extras" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, }),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Match.V1.SpectatorBattleEnd), global::Thetan.Match.V1.SpectatorBattleEnd.Parser, new[]{ "MatchID", "Rounds", "Players", "EndAt" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class MatchRound : pb::IMessage<MatchRound>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MatchRound> _parser = new pb::MessageParser<MatchRound>(() => new MatchRound());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<MatchRound> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Match.V1.MatchRivalsReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchRound() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchRound(MatchRound other) : this() {
      playerInRounds_ = other.playerInRounds_.Clone();
      gameInputs_ = other.gameInputs_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchRound Clone() {
      return new MatchRound(this);
    }

    /// <summary>Field number for the "playerInRounds" field.</summary>
    public const int PlayerInRoundsFieldNumber = 1;
    private static readonly pb::FieldCodec<string> _repeated_playerInRounds_codec
        = pb::FieldCodec.ForString(10);
    private readonly pbc::RepeatedField<string> playerInRounds_ = new pbc::RepeatedField<string>();
    /// <summary>
    /// List of players played in this round, playerID (ObjectID)
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<string> PlayerInRounds {
      get { return playerInRounds_; }
    }

    /// <summary>Field number for the "gameInputs" field.</summary>
    public const int GameInputsFieldNumber = 2;
    private pb::ByteString gameInputs_ = pb::ByteString.Empty;
    /// <summary>
    /// Game inputs contains Message Pack (compressed) data for all players in this round.
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pb::ByteString GameInputs {
      get { return gameInputs_; }
      set {
        gameInputs_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as MatchRound);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(MatchRound other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!playerInRounds_.Equals(other.playerInRounds_)) return false;
      if (GameInputs != other.GameInputs) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= playerInRounds_.GetHashCode();
      if (GameInputs.Length != 0) hash ^= GameInputs.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      playerInRounds_.WriteTo(output, _repeated_playerInRounds_codec);
      if (GameInputs.Length != 0) {
        output.WriteRawTag(18);
        output.WriteBytes(GameInputs);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      playerInRounds_.WriteTo(ref output, _repeated_playerInRounds_codec);
      if (GameInputs.Length != 0) {
        output.WriteRawTag(18);
        output.WriteBytes(GameInputs);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      size += playerInRounds_.CalculateSize(_repeated_playerInRounds_codec);
      if (GameInputs.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeBytesSize(GameInputs);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(MatchRound other) {
      if (other == null) {
        return;
      }
      playerInRounds_.Add(other.playerInRounds_);
      if (other.GameInputs.Length != 0) {
        GameInputs = other.GameInputs;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            playerInRounds_.AddEntriesFrom(input, _repeated_playerInRounds_codec);
            break;
          }
          case 18: {
            GameInputs = input.ReadBytes();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            playerInRounds_.AddEntriesFrom(ref input, _repeated_playerInRounds_codec);
            break;
          }
          case 18: {
            GameInputs = input.ReadBytes();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class MatchPlayer : pb::IMessage<MatchPlayer>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MatchPlayer> _parser = new pb::MessageParser<MatchPlayer>(() => new MatchPlayer());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<MatchPlayer> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Match.V1.MatchRivalsReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchPlayer() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchPlayer(MatchPlayer other) : this() {
      playerID_ = other.playerID_;
      battleRank_ = other.battleRank_;
      roundTimes_ = other.roundTimes_.Clone();
      extras_ = other.extras_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchPlayer Clone() {
      return new MatchPlayer(this);
    }

    /// <summary>Field number for the "playerID" field.</summary>
    public const int PlayerIDFieldNumber = 1;
    private string playerID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string PlayerID {
      get { return playerID_; }
      set {
        playerID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "battleRank" field.</summary>
    public const int BattleRankFieldNumber = 2;
    private int battleRank_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int BattleRank {
      get { return battleRank_; }
      set {
        battleRank_ = value;
      }
    }

    /// <summary>Field number for the "roundTimes" field.</summary>
    public const int RoundTimesFieldNumber = 3;
    private static readonly pb::FieldCodec<float> _repeated_roundTimes_codec
        = pb::FieldCodec.ForFloat(26);
    private readonly pbc::RepeatedField<float> roundTimes_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<float> RoundTimes {
      get { return roundTimes_; }
    }

    /// <summary>Field number for the "extras" field.</summary>
    public const int ExtrasFieldNumber = 4;
    private static readonly pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any>.Codec _map_extras_codec
        = new pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any>.Codec(pb::FieldCodec.ForString(10, ""), pb::FieldCodec.ForMessage(18, global::Google.Protobuf.WellKnownTypes.Any.Parser), 34);
    private readonly pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any> extras_ = new pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any> Extras {
      get { return extras_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as MatchPlayer);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(MatchPlayer other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (PlayerID != other.PlayerID) return false;
      if (BattleRank != other.BattleRank) return false;
      if(!roundTimes_.Equals(other.roundTimes_)) return false;
      if (!Extras.Equals(other.Extras)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (PlayerID.Length != 0) hash ^= PlayerID.GetHashCode();
      if (BattleRank != 0) hash ^= BattleRank.GetHashCode();
      hash ^= roundTimes_.GetHashCode();
      hash ^= Extras.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (PlayerID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerID);
      }
      if (BattleRank != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(BattleRank);
      }
      roundTimes_.WriteTo(output, _repeated_roundTimes_codec);
      extras_.WriteTo(output, _map_extras_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (PlayerID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerID);
      }
      if (BattleRank != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(BattleRank);
      }
      roundTimes_.WriteTo(ref output, _repeated_roundTimes_codec);
      extras_.WriteTo(ref output, _map_extras_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (PlayerID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(PlayerID);
      }
      if (BattleRank != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(BattleRank);
      }
      size += roundTimes_.CalculateSize(_repeated_roundTimes_codec);
      size += extras_.CalculateSize(_map_extras_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(MatchPlayer other) {
      if (other == null) {
        return;
      }
      if (other.PlayerID.Length != 0) {
        PlayerID = other.PlayerID;
      }
      if (other.BattleRank != 0) {
        BattleRank = other.BattleRank;
      }
      roundTimes_.Add(other.roundTimes_);
      extras_.Add(other.extras_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            PlayerID = input.ReadString();
            break;
          }
          case 16: {
            BattleRank = input.ReadInt32();
            break;
          }
          case 26:
          case 29: {
            roundTimes_.AddEntriesFrom(input, _repeated_roundTimes_codec);
            break;
          }
          case 34: {
            extras_.AddEntriesFrom(input, _map_extras_codec);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            PlayerID = input.ReadString();
            break;
          }
          case 16: {
            BattleRank = input.ReadInt32();
            break;
          }
          case 26:
          case 29: {
            roundTimes_.AddEntriesFrom(ref input, _repeated_roundTimes_codec);
            break;
          }
          case 34: {
            extras_.AddEntriesFrom(ref input, _map_extras_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class SpectatorBattleEnd : pb::IMessage<SpectatorBattleEnd>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<SpectatorBattleEnd> _parser = new pb::MessageParser<SpectatorBattleEnd>(() => new SpectatorBattleEnd());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<SpectatorBattleEnd> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Match.V1.MatchRivalsReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SpectatorBattleEnd() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SpectatorBattleEnd(SpectatorBattleEnd other) : this() {
      matchID_ = other.matchID_;
      rounds_ = other.rounds_.Clone();
      players_ = other.players_.Clone();
      endAt_ = other.endAt_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SpectatorBattleEnd Clone() {
      return new SpectatorBattleEnd(this);
    }

    /// <summary>Field number for the "matchID" field.</summary>
    public const int MatchIDFieldNumber = 1;
    private string matchID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string MatchID {
      get { return matchID_; }
      set {
        matchID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "rounds" field.</summary>
    public const int RoundsFieldNumber = 2;
    private static readonly pb::FieldCodec<global::Thetan.Match.V1.MatchRound> _repeated_rounds_codec
        = pb::FieldCodec.ForMessage(18, global::Thetan.Match.V1.MatchRound.Parser);
    private readonly pbc::RepeatedField<global::Thetan.Match.V1.MatchRound> rounds_ = new pbc::RepeatedField<global::Thetan.Match.V1.MatchRound>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<global::Thetan.Match.V1.MatchRound> Rounds {
      get { return rounds_; }
    }

    /// <summary>Field number for the "players" field.</summary>
    public const int PlayersFieldNumber = 3;
    private static readonly pb::FieldCodec<global::Thetan.Match.V1.MatchPlayer> _repeated_players_codec
        = pb::FieldCodec.ForMessage(26, global::Thetan.Match.V1.MatchPlayer.Parser);
    private readonly pbc::RepeatedField<global::Thetan.Match.V1.MatchPlayer> players_ = new pbc::RepeatedField<global::Thetan.Match.V1.MatchPlayer>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<global::Thetan.Match.V1.MatchPlayer> Players {
      get { return players_; }
    }

    /// <summary>Field number for the "endAt" field.</summary>
    public const int EndAtFieldNumber = 4;
    private long endAt_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long EndAt {
      get { return endAt_; }
      set {
        endAt_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as SpectatorBattleEnd);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(SpectatorBattleEnd other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (MatchID != other.MatchID) return false;
      if(!rounds_.Equals(other.rounds_)) return false;
      if(!players_.Equals(other.players_)) return false;
      if (EndAt != other.EndAt) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (MatchID.Length != 0) hash ^= MatchID.GetHashCode();
      hash ^= rounds_.GetHashCode();
      hash ^= players_.GetHashCode();
      if (EndAt != 0L) hash ^= EndAt.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (MatchID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(MatchID);
      }
      rounds_.WriteTo(output, _repeated_rounds_codec);
      players_.WriteTo(output, _repeated_players_codec);
      if (EndAt != 0L) {
        output.WriteRawTag(32);
        output.WriteInt64(EndAt);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (MatchID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(MatchID);
      }
      rounds_.WriteTo(ref output, _repeated_rounds_codec);
      players_.WriteTo(ref output, _repeated_players_codec);
      if (EndAt != 0L) {
        output.WriteRawTag(32);
        output.WriteInt64(EndAt);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (MatchID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(MatchID);
      }
      size += rounds_.CalculateSize(_repeated_rounds_codec);
      size += players_.CalculateSize(_repeated_players_codec);
      if (EndAt != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(EndAt);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(SpectatorBattleEnd other) {
      if (other == null) {
        return;
      }
      if (other.MatchID.Length != 0) {
        MatchID = other.MatchID;
      }
      rounds_.Add(other.rounds_);
      players_.Add(other.players_);
      if (other.EndAt != 0L) {
        EndAt = other.EndAt;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            MatchID = input.ReadString();
            break;
          }
          case 18: {
            rounds_.AddEntriesFrom(input, _repeated_rounds_codec);
            break;
          }
          case 26: {
            players_.AddEntriesFrom(input, _repeated_players_codec);
            break;
          }
          case 32: {
            EndAt = input.ReadInt64();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            MatchID = input.ReadString();
            break;
          }
          case 18: {
            rounds_.AddEntriesFrom(ref input, _repeated_rounds_codec);
            break;
          }
          case 26: {
            players_.AddEntriesFrom(ref input, _repeated_players_codec);
            break;
          }
          case 32: {
            EndAt = input.ReadInt64();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
