// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: mmf-service.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Mmf {

  /// <summary>Holder for reflection information generated from mmf-service.proto</summary>
  public static partial class MmfServiceReflection {

    #region Descriptor
    /// <summary>File descriptor for mmf-service.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MmfServiceReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChFtbWYtc2VydmljZS5wcm90bxIDbW1mIt0BCg1NbWZUaWNrZXREYXRhEiAK",
            "C2NvdW50UGxheWVyGAEgASgFUgtjb3VudFBsYXllchIkCg1iZWhhdmlvclBv",
            "aW50GAIgASgFUg1iZWhhdmlvclBvaW50EhwKCWlzSGVyb05GVBgDIAEoCFIJ",
            "aXNIZXJvTkZUEiIKDHRyb3BoaWVzUmFuaxgEIAEoBVIMdHJvcGhpZXNSYW5r",
            "EhgKB3JlZ2lvbnMYBSADKAVSB3JlZ2lvbnMSKAoPbWF0Y2hTZWFyY2hUeXBl",
            "GAYgASgFUg9tYXRjaFNlYXJjaFR5cGVCVAoHY29tLm1tZkIPTW1mU2Vydmlj",
            "ZVByb3RvSAFQAVoKL2NvcmVwcm90b6ICA01YWKoCA01tZsoCA01tZuICD01t",
            "ZlxHUEJNZXRhZGF0YeoCA01tZmIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Mmf.MmfTicketData), global::Mmf.MmfTicketData.Parser, new[]{ "CountPlayer", "BehaviorPoint", "IsHeroNFT", "TrophiesRank", "Regions", "MatchSearchType" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class MmfTicketData : pb::IMessage<MmfTicketData>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MmfTicketData> _parser = new pb::MessageParser<MmfTicketData>(() => new MmfTicketData());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<MmfTicketData> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Mmf.MmfServiceReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MmfTicketData() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MmfTicketData(MmfTicketData other) : this() {
      countPlayer_ = other.countPlayer_;
      behaviorPoint_ = other.behaviorPoint_;
      isHeroNFT_ = other.isHeroNFT_;
      trophiesRank_ = other.trophiesRank_;
      regions_ = other.regions_.Clone();
      matchSearchType_ = other.matchSearchType_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MmfTicketData Clone() {
      return new MmfTicketData(this);
    }

    /// <summary>Field number for the "countPlayer" field.</summary>
    public const int CountPlayerFieldNumber = 1;
    private int countPlayer_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CountPlayer {
      get { return countPlayer_; }
      set {
        countPlayer_ = value;
      }
    }

    /// <summary>Field number for the "behaviorPoint" field.</summary>
    public const int BehaviorPointFieldNumber = 2;
    private int behaviorPoint_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int BehaviorPoint {
      get { return behaviorPoint_; }
      set {
        behaviorPoint_ = value;
      }
    }

    /// <summary>Field number for the "isHeroNFT" field.</summary>
    public const int IsHeroNFTFieldNumber = 3;
    private bool isHeroNFT_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool IsHeroNFT {
      get { return isHeroNFT_; }
      set {
        isHeroNFT_ = value;
      }
    }

    /// <summary>Field number for the "trophiesRank" field.</summary>
    public const int TrophiesRankFieldNumber = 4;
    private int trophiesRank_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int TrophiesRank {
      get { return trophiesRank_; }
      set {
        trophiesRank_ = value;
      }
    }

    /// <summary>Field number for the "regions" field.</summary>
    public const int RegionsFieldNumber = 5;
    private static readonly pb::FieldCodec<int> _repeated_regions_codec
        = pb::FieldCodec.ForInt32(42);
    private readonly pbc::RepeatedField<int> regions_ = new pbc::RepeatedField<int>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<int> Regions {
      get { return regions_; }
    }

    /// <summary>Field number for the "matchSearchType" field.</summary>
    public const int MatchSearchTypeFieldNumber = 6;
    private int matchSearchType_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int MatchSearchType {
      get { return matchSearchType_; }
      set {
        matchSearchType_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as MmfTicketData);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(MmfTicketData other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (CountPlayer != other.CountPlayer) return false;
      if (BehaviorPoint != other.BehaviorPoint) return false;
      if (IsHeroNFT != other.IsHeroNFT) return false;
      if (TrophiesRank != other.TrophiesRank) return false;
      if(!regions_.Equals(other.regions_)) return false;
      if (MatchSearchType != other.MatchSearchType) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (CountPlayer != 0) hash ^= CountPlayer.GetHashCode();
      if (BehaviorPoint != 0) hash ^= BehaviorPoint.GetHashCode();
      if (IsHeroNFT != false) hash ^= IsHeroNFT.GetHashCode();
      if (TrophiesRank != 0) hash ^= TrophiesRank.GetHashCode();
      hash ^= regions_.GetHashCode();
      if (MatchSearchType != 0) hash ^= MatchSearchType.GetHashCode();
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
      if (CountPlayer != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(CountPlayer);
      }
      if (BehaviorPoint != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(BehaviorPoint);
      }
      if (IsHeroNFT != false) {
        output.WriteRawTag(24);
        output.WriteBool(IsHeroNFT);
      }
      if (TrophiesRank != 0) {
        output.WriteRawTag(32);
        output.WriteInt32(TrophiesRank);
      }
      regions_.WriteTo(output, _repeated_regions_codec);
      if (MatchSearchType != 0) {
        output.WriteRawTag(48);
        output.WriteInt32(MatchSearchType);
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
      if (CountPlayer != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(CountPlayer);
      }
      if (BehaviorPoint != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(BehaviorPoint);
      }
      if (IsHeroNFT != false) {
        output.WriteRawTag(24);
        output.WriteBool(IsHeroNFT);
      }
      if (TrophiesRank != 0) {
        output.WriteRawTag(32);
        output.WriteInt32(TrophiesRank);
      }
      regions_.WriteTo(ref output, _repeated_regions_codec);
      if (MatchSearchType != 0) {
        output.WriteRawTag(48);
        output.WriteInt32(MatchSearchType);
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
      if (CountPlayer != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(CountPlayer);
      }
      if (BehaviorPoint != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(BehaviorPoint);
      }
      if (IsHeroNFT != false) {
        size += 1 + 1;
      }
      if (TrophiesRank != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(TrophiesRank);
      }
      size += regions_.CalculateSize(_repeated_regions_codec);
      if (MatchSearchType != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(MatchSearchType);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(MmfTicketData other) {
      if (other == null) {
        return;
      }
      if (other.CountPlayer != 0) {
        CountPlayer = other.CountPlayer;
      }
      if (other.BehaviorPoint != 0) {
        BehaviorPoint = other.BehaviorPoint;
      }
      if (other.IsHeroNFT != false) {
        IsHeroNFT = other.IsHeroNFT;
      }
      if (other.TrophiesRank != 0) {
        TrophiesRank = other.TrophiesRank;
      }
      regions_.Add(other.regions_);
      if (other.MatchSearchType != 0) {
        MatchSearchType = other.MatchSearchType;
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
            CountPlayer = input.ReadInt32();
            break;
          }
          case 16: {
            BehaviorPoint = input.ReadInt32();
            break;
          }
          case 24: {
            IsHeroNFT = input.ReadBool();
            break;
          }
          case 32: {
            TrophiesRank = input.ReadInt32();
            break;
          }
          case 42:
          case 40: {
            regions_.AddEntriesFrom(input, _repeated_regions_codec);
            break;
          }
          case 48: {
            MatchSearchType = input.ReadInt32();
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
            CountPlayer = input.ReadInt32();
            break;
          }
          case 16: {
            BehaviorPoint = input.ReadInt32();
            break;
          }
          case 24: {
            IsHeroNFT = input.ReadBool();
            break;
          }
          case 32: {
            TrophiesRank = input.ReadInt32();
            break;
          }
          case 42:
          case 40: {
            regions_.AddEntriesFrom(ref input, _repeated_regions_codec);
            break;
          }
          case 48: {
            MatchSearchType = input.ReadInt32();
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
