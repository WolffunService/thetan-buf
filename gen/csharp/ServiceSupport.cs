// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/support/v1/service_support.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Support.V1 {

  /// <summary>Holder for reflection information generated from thetan/support/v1/service_support.proto</summary>
  public static partial class ServiceSupportReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/support/v1/service_support.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ServiceSupportReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cid0aGV0YW4vc3VwcG9ydC92MS9zZXJ2aWNlX3N1cHBvcnQucHJvdG8SEXRo",
            "ZXRhbi5zdXBwb3J0LnYxIsICChFTZWFyY2hCb3RzUmVxdWVzdBIxCgZnYW1l",
            "SUQYASABKA4yGS50aGV0YW4uc3VwcG9ydC52MS5HYW1lSURSBmdhbWVJRBIS",
            "CgRyYW5rGAIgASgFUgRyYW5rEhoKCHJlZ2lvbklEGAMgASgFUghyZWdpb25J",
            "RBI0CgdmZWF0dXJlGAQgASgOMhoudGhldGFuLnN1cHBvcnQudjEuRmVhdHVy",
            "ZVIHZmVhdHVyZRIeCgphcnJDb3VudHJ5GAUgAygJUgphcnJDb3VudHJ5EiIK",
            "DHVzZVRpbWVJblNlYxgGIAEoA1IMdXNlVGltZUluU2VjEiIKDGJhdHRsZU51",
            "bWJlchgHIAEoBVIMYmF0dGxlTnVtYmVyEhgKB251bWJlcnMYCCABKAVSB251",
            "bWJlcnMSEgoEc2tpcBgJIAEoBVIEc2tpcCJEChJTZWFyY2hCb3RzUmVzcG9u",
            "c2USLgoEYm90cxgBIAMoCzIaLnRoZXRhbi5zdXBwb3J0LnYxLkJvdEluZm9S",
            "BGJvdHMiiwEKB0JvdEluZm8SFgoGdXNlcklEGAEgASgJUgZ1c2VySUQSGgoI",
            "dXNlcm5hbWUYAiABKAlSCHVzZXJuYW1lEiAKC2NvdW50cnlDb2RlGAMgASgJ",
            "Ugtjb3VudHJ5Q29kZRISCgRyYW5rGAQgASgFUgRyYW5rEhYKBnRyb3BoeRgF",
            "IAEoBVIGdHJvcGh5KlYKBkdhbWVJRBIQCgxHQU1FX0lEX05PTkUQABIRCg1H",
            "QU1FX0lEX0FSRU5BEAESEQoNR0FNRV9JRF9SSVZBTBACEhQKEEdBTUVfSURf",
            "SU1NT1JUQUwQAyprCgdGZWF0dXJlEhAKDEZFQVRVUkVfTk9ORRAAEhMKD0ZF",
            "QVRVUkVfSU5fR0FNRRABEhEKDUZFQVRVUkVfTE9CQlkQAhIOCgpGRUFUVVJF",
            "X0xCEAMSFgoSRkVBVFVSRV9DSEFUX1NISUxMEAQybQoOU3VwcG9ydFNlcnZp",
            "Y2USWwoKU2VhcmNoQm90cxIkLnRoZXRhbi5zdXBwb3J0LnYxLlNlYXJjaEJv",
            "dHNSZXF1ZXN0GiUudGhldGFuLnN1cHBvcnQudjEuU2VhcmNoQm90c1Jlc3Bv",
            "bnNlIgBCyQEKFWNvbS50aGV0YW4uc3VwcG9ydC52MUITU2VydmljZVN1cHBv",
            "cnRQcm90b1ABWjV0aGV0YW4tYnVmL2dlbi9nby90aGV0YW4vc3VwcG9ydC92",
            "MTt0aGV0YW5fc3VwcG9ydF92MaICA1RTWKoCEVRoZXRhbi5TdXBwb3J0LlYx",
            "ygIRVGhldGFuXFN1cHBvcnRcVjHiAh1UaGV0YW5cU3VwcG9ydFxWMVxHUEJN",
            "ZXRhZGF0YeoCE1RoZXRhbjo6U3VwcG9ydDo6VjFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Thetan.Support.V1.GameID), typeof(global::Thetan.Support.V1.Feature), }, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Support.V1.SearchBotsRequest), global::Thetan.Support.V1.SearchBotsRequest.Parser, new[]{ "GameID", "Rank", "RegionID", "Feature", "ArrCountry", "UseTimeInSec", "BattleNumber", "Numbers", "Skip" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Support.V1.SearchBotsResponse), global::Thetan.Support.V1.SearchBotsResponse.Parser, new[]{ "Bots" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Support.V1.BotInfo), global::Thetan.Support.V1.BotInfo.Parser, new[]{ "UserID", "Username", "CountryCode", "Rank", "Trophy" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  public enum GameID {
    [pbr::OriginalName("GAME_ID_NONE")] None = 0,
    [pbr::OriginalName("GAME_ID_ARENA")] Arena = 1,
    [pbr::OriginalName("GAME_ID_RIVAL")] Rival = 2,
    [pbr::OriginalName("GAME_ID_IMMORTAL")] Immortal = 3,
  }

  public enum Feature {
    [pbr::OriginalName("FEATURE_NONE")] None = 0,
    [pbr::OriginalName("FEATURE_IN_GAME")] InGame = 1,
    [pbr::OriginalName("FEATURE_LOBBY")] Lobby = 2,
    [pbr::OriginalName("FEATURE_LB")] Lb = 3,
    [pbr::OriginalName("FEATURE_CHAT_SHILL")] ChatShill = 4,
  }

  #endregion

  #region Messages
  public sealed partial class SearchBotsRequest : pb::IMessage<SearchBotsRequest>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<SearchBotsRequest> _parser = new pb::MessageParser<SearchBotsRequest>(() => new SearchBotsRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<SearchBotsRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Support.V1.ServiceSupportReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SearchBotsRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SearchBotsRequest(SearchBotsRequest other) : this() {
      gameID_ = other.gameID_;
      rank_ = other.rank_;
      regionID_ = other.regionID_;
      feature_ = other.feature_;
      arrCountry_ = other.arrCountry_.Clone();
      useTimeInSec_ = other.useTimeInSec_;
      battleNumber_ = other.battleNumber_;
      numbers_ = other.numbers_;
      skip_ = other.skip_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SearchBotsRequest Clone() {
      return new SearchBotsRequest(this);
    }

    /// <summary>Field number for the "gameID" field.</summary>
    public const int GameIDFieldNumber = 1;
    private global::Thetan.Support.V1.GameID gameID_ = global::Thetan.Support.V1.GameID.None;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Support.V1.GameID GameID {
      get { return gameID_; }
      set {
        gameID_ = value;
      }
    }

    /// <summary>Field number for the "rank" field.</summary>
    public const int RankFieldNumber = 2;
    private int rank_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Rank {
      get { return rank_; }
      set {
        rank_ = value;
      }
    }

    /// <summary>Field number for the "regionID" field.</summary>
    public const int RegionIDFieldNumber = 3;
    private int regionID_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int RegionID {
      get { return regionID_; }
      set {
        regionID_ = value;
      }
    }

    /// <summary>Field number for the "feature" field.</summary>
    public const int FeatureFieldNumber = 4;
    private global::Thetan.Support.V1.Feature feature_ = global::Thetan.Support.V1.Feature.None;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Support.V1.Feature Feature {
      get { return feature_; }
      set {
        feature_ = value;
      }
    }

    /// <summary>Field number for the "arrCountry" field.</summary>
    public const int ArrCountryFieldNumber = 5;
    private static readonly pb::FieldCodec<string> _repeated_arrCountry_codec
        = pb::FieldCodec.ForString(42);
    private readonly pbc::RepeatedField<string> arrCountry_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<string> ArrCountry {
      get { return arrCountry_; }
    }

    /// <summary>Field number for the "useTimeInSec" field.</summary>
    public const int UseTimeInSecFieldNumber = 6;
    private long useTimeInSec_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long UseTimeInSec {
      get { return useTimeInSec_; }
      set {
        useTimeInSec_ = value;
      }
    }

    /// <summary>Field number for the "battleNumber" field.</summary>
    public const int BattleNumberFieldNumber = 7;
    private int battleNumber_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int BattleNumber {
      get { return battleNumber_; }
      set {
        battleNumber_ = value;
      }
    }

    /// <summary>Field number for the "numbers" field.</summary>
    public const int NumbersFieldNumber = 8;
    private int numbers_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Numbers {
      get { return numbers_; }
      set {
        numbers_ = value;
      }
    }

    /// <summary>Field number for the "skip" field.</summary>
    public const int SkipFieldNumber = 9;
    private int skip_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Skip {
      get { return skip_; }
      set {
        skip_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as SearchBotsRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(SearchBotsRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (GameID != other.GameID) return false;
      if (Rank != other.Rank) return false;
      if (RegionID != other.RegionID) return false;
      if (Feature != other.Feature) return false;
      if(!arrCountry_.Equals(other.arrCountry_)) return false;
      if (UseTimeInSec != other.UseTimeInSec) return false;
      if (BattleNumber != other.BattleNumber) return false;
      if (Numbers != other.Numbers) return false;
      if (Skip != other.Skip) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (GameID != global::Thetan.Support.V1.GameID.None) hash ^= GameID.GetHashCode();
      if (Rank != 0) hash ^= Rank.GetHashCode();
      if (RegionID != 0) hash ^= RegionID.GetHashCode();
      if (Feature != global::Thetan.Support.V1.Feature.None) hash ^= Feature.GetHashCode();
      hash ^= arrCountry_.GetHashCode();
      if (UseTimeInSec != 0L) hash ^= UseTimeInSec.GetHashCode();
      if (BattleNumber != 0) hash ^= BattleNumber.GetHashCode();
      if (Numbers != 0) hash ^= Numbers.GetHashCode();
      if (Skip != 0) hash ^= Skip.GetHashCode();
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
      if (GameID != global::Thetan.Support.V1.GameID.None) {
        output.WriteRawTag(8);
        output.WriteEnum((int) GameID);
      }
      if (Rank != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Rank);
      }
      if (RegionID != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(RegionID);
      }
      if (Feature != global::Thetan.Support.V1.Feature.None) {
        output.WriteRawTag(32);
        output.WriteEnum((int) Feature);
      }
      arrCountry_.WriteTo(output, _repeated_arrCountry_codec);
      if (UseTimeInSec != 0L) {
        output.WriteRawTag(48);
        output.WriteInt64(UseTimeInSec);
      }
      if (BattleNumber != 0) {
        output.WriteRawTag(56);
        output.WriteInt32(BattleNumber);
      }
      if (Numbers != 0) {
        output.WriteRawTag(64);
        output.WriteInt32(Numbers);
      }
      if (Skip != 0) {
        output.WriteRawTag(72);
        output.WriteInt32(Skip);
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
      if (GameID != global::Thetan.Support.V1.GameID.None) {
        output.WriteRawTag(8);
        output.WriteEnum((int) GameID);
      }
      if (Rank != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Rank);
      }
      if (RegionID != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(RegionID);
      }
      if (Feature != global::Thetan.Support.V1.Feature.None) {
        output.WriteRawTag(32);
        output.WriteEnum((int) Feature);
      }
      arrCountry_.WriteTo(ref output, _repeated_arrCountry_codec);
      if (UseTimeInSec != 0L) {
        output.WriteRawTag(48);
        output.WriteInt64(UseTimeInSec);
      }
      if (BattleNumber != 0) {
        output.WriteRawTag(56);
        output.WriteInt32(BattleNumber);
      }
      if (Numbers != 0) {
        output.WriteRawTag(64);
        output.WriteInt32(Numbers);
      }
      if (Skip != 0) {
        output.WriteRawTag(72);
        output.WriteInt32(Skip);
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
      if (GameID != global::Thetan.Support.V1.GameID.None) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) GameID);
      }
      if (Rank != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Rank);
      }
      if (RegionID != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(RegionID);
      }
      if (Feature != global::Thetan.Support.V1.Feature.None) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Feature);
      }
      size += arrCountry_.CalculateSize(_repeated_arrCountry_codec);
      if (UseTimeInSec != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(UseTimeInSec);
      }
      if (BattleNumber != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(BattleNumber);
      }
      if (Numbers != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Numbers);
      }
      if (Skip != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Skip);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(SearchBotsRequest other) {
      if (other == null) {
        return;
      }
      if (other.GameID != global::Thetan.Support.V1.GameID.None) {
        GameID = other.GameID;
      }
      if (other.Rank != 0) {
        Rank = other.Rank;
      }
      if (other.RegionID != 0) {
        RegionID = other.RegionID;
      }
      if (other.Feature != global::Thetan.Support.V1.Feature.None) {
        Feature = other.Feature;
      }
      arrCountry_.Add(other.arrCountry_);
      if (other.UseTimeInSec != 0L) {
        UseTimeInSec = other.UseTimeInSec;
      }
      if (other.BattleNumber != 0) {
        BattleNumber = other.BattleNumber;
      }
      if (other.Numbers != 0) {
        Numbers = other.Numbers;
      }
      if (other.Skip != 0) {
        Skip = other.Skip;
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
          case 8: {
            GameID = (global::Thetan.Support.V1.GameID) input.ReadEnum();
            break;
          }
          case 16: {
            Rank = input.ReadInt32();
            break;
          }
          case 24: {
            RegionID = input.ReadInt32();
            break;
          }
          case 32: {
            Feature = (global::Thetan.Support.V1.Feature) input.ReadEnum();
            break;
          }
          case 42: {
            arrCountry_.AddEntriesFrom(input, _repeated_arrCountry_codec);
            break;
          }
          case 48: {
            UseTimeInSec = input.ReadInt64();
            break;
          }
          case 56: {
            BattleNumber = input.ReadInt32();
            break;
          }
          case 64: {
            Numbers = input.ReadInt32();
            break;
          }
          case 72: {
            Skip = input.ReadInt32();
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
          case 8: {
            GameID = (global::Thetan.Support.V1.GameID) input.ReadEnum();
            break;
          }
          case 16: {
            Rank = input.ReadInt32();
            break;
          }
          case 24: {
            RegionID = input.ReadInt32();
            break;
          }
          case 32: {
            Feature = (global::Thetan.Support.V1.Feature) input.ReadEnum();
            break;
          }
          case 42: {
            arrCountry_.AddEntriesFrom(ref input, _repeated_arrCountry_codec);
            break;
          }
          case 48: {
            UseTimeInSec = input.ReadInt64();
            break;
          }
          case 56: {
            BattleNumber = input.ReadInt32();
            break;
          }
          case 64: {
            Numbers = input.ReadInt32();
            break;
          }
          case 72: {
            Skip = input.ReadInt32();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class SearchBotsResponse : pb::IMessage<SearchBotsResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<SearchBotsResponse> _parser = new pb::MessageParser<SearchBotsResponse>(() => new SearchBotsResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<SearchBotsResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Support.V1.ServiceSupportReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SearchBotsResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SearchBotsResponse(SearchBotsResponse other) : this() {
      bots_ = other.bots_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SearchBotsResponse Clone() {
      return new SearchBotsResponse(this);
    }

    /// <summary>Field number for the "bots" field.</summary>
    public const int BotsFieldNumber = 1;
    private static readonly pb::FieldCodec<global::Thetan.Support.V1.BotInfo> _repeated_bots_codec
        = pb::FieldCodec.ForMessage(10, global::Thetan.Support.V1.BotInfo.Parser);
    private readonly pbc::RepeatedField<global::Thetan.Support.V1.BotInfo> bots_ = new pbc::RepeatedField<global::Thetan.Support.V1.BotInfo>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<global::Thetan.Support.V1.BotInfo> Bots {
      get { return bots_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as SearchBotsResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(SearchBotsResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!bots_.Equals(other.bots_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= bots_.GetHashCode();
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
      bots_.WriteTo(output, _repeated_bots_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      bots_.WriteTo(ref output, _repeated_bots_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      size += bots_.CalculateSize(_repeated_bots_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(SearchBotsResponse other) {
      if (other == null) {
        return;
      }
      bots_.Add(other.bots_);
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
            bots_.AddEntriesFrom(input, _repeated_bots_codec);
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
            bots_.AddEntriesFrom(ref input, _repeated_bots_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class BotInfo : pb::IMessage<BotInfo>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<BotInfo> _parser = new pb::MessageParser<BotInfo>(() => new BotInfo());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<BotInfo> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Support.V1.ServiceSupportReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public BotInfo() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public BotInfo(BotInfo other) : this() {
      userID_ = other.userID_;
      username_ = other.username_;
      countryCode_ = other.countryCode_;
      rank_ = other.rank_;
      trophy_ = other.trophy_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public BotInfo Clone() {
      return new BotInfo(this);
    }

    /// <summary>Field number for the "userID" field.</summary>
    public const int UserIDFieldNumber = 1;
    private string userID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string UserID {
      get { return userID_; }
      set {
        userID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "username" field.</summary>
    public const int UsernameFieldNumber = 2;
    private string username_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Username {
      get { return username_; }
      set {
        username_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "countryCode" field.</summary>
    public const int CountryCodeFieldNumber = 3;
    private string countryCode_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string CountryCode {
      get { return countryCode_; }
      set {
        countryCode_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "rank" field.</summary>
    public const int RankFieldNumber = 4;
    private int rank_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Rank {
      get { return rank_; }
      set {
        rank_ = value;
      }
    }

    /// <summary>Field number for the "trophy" field.</summary>
    public const int TrophyFieldNumber = 5;
    private int trophy_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Trophy {
      get { return trophy_; }
      set {
        trophy_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as BotInfo);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(BotInfo other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (UserID != other.UserID) return false;
      if (Username != other.Username) return false;
      if (CountryCode != other.CountryCode) return false;
      if (Rank != other.Rank) return false;
      if (Trophy != other.Trophy) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (UserID.Length != 0) hash ^= UserID.GetHashCode();
      if (Username.Length != 0) hash ^= Username.GetHashCode();
      if (CountryCode.Length != 0) hash ^= CountryCode.GetHashCode();
      if (Rank != 0) hash ^= Rank.GetHashCode();
      if (Trophy != 0) hash ^= Trophy.GetHashCode();
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
      if (UserID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(UserID);
      }
      if (Username.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Username);
      }
      if (CountryCode.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(CountryCode);
      }
      if (Rank != 0) {
        output.WriteRawTag(32);
        output.WriteInt32(Rank);
      }
      if (Trophy != 0) {
        output.WriteRawTag(40);
        output.WriteInt32(Trophy);
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
      if (UserID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(UserID);
      }
      if (Username.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Username);
      }
      if (CountryCode.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(CountryCode);
      }
      if (Rank != 0) {
        output.WriteRawTag(32);
        output.WriteInt32(Rank);
      }
      if (Trophy != 0) {
        output.WriteRawTag(40);
        output.WriteInt32(Trophy);
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
      if (UserID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(UserID);
      }
      if (Username.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Username);
      }
      if (CountryCode.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(CountryCode);
      }
      if (Rank != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Rank);
      }
      if (Trophy != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Trophy);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(BotInfo other) {
      if (other == null) {
        return;
      }
      if (other.UserID.Length != 0) {
        UserID = other.UserID;
      }
      if (other.Username.Length != 0) {
        Username = other.Username;
      }
      if (other.CountryCode.Length != 0) {
        CountryCode = other.CountryCode;
      }
      if (other.Rank != 0) {
        Rank = other.Rank;
      }
      if (other.Trophy != 0) {
        Trophy = other.Trophy;
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
            UserID = input.ReadString();
            break;
          }
          case 18: {
            Username = input.ReadString();
            break;
          }
          case 26: {
            CountryCode = input.ReadString();
            break;
          }
          case 32: {
            Rank = input.ReadInt32();
            break;
          }
          case 40: {
            Trophy = input.ReadInt32();
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
            UserID = input.ReadString();
            break;
          }
          case 18: {
            Username = input.ReadString();
            break;
          }
          case 26: {
            CountryCode = input.ReadString();
            break;
          }
          case 32: {
            Rank = input.ReadInt32();
            break;
          }
          case 40: {
            Trophy = input.ReadInt32();
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
