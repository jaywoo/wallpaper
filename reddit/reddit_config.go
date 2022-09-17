package reddit

// IndexDomain 首页域名
const IndexDomain = "https://www.reddit.com"

//IndexURL 首页URI
const IndexURL = "/r/wallpapers.json"

const PageDomain = "https://gateway.reddit.com"
const PageURL = "/desktopapi/v1/subreddits/wallpapers?rtj=only&redditWebClient=web2x&app=web2x-client-production&include=prefsSubreddit&after=t3_xaj8zg&dist=3&forceGeopopular=false&layout=card&sort=hot"

type RedditPageData struct {
    Posts                map[string]PagePost `json:"posts"`
    Token                string              `json:"token"`
    Dist                 int                 `json:"dist"`
    Account              interface{}         `json:"account"`
    Features             interface{}         `json:"features"`
    ListingSort          string              `json:"listingSort"`
    SubredditPermissions interface{}         `json:"subredditPermissions"`
    Preferences          interface{}         `json:"preferences"`
    StructuredStyles     interface{}         `json:"structuredStyles"`
}
type PagePost struct {
    ID              string      `json:"id"`
    NumComments     int         `json:"numComments"`
    Created         int64       `json:"created"`
    Score           int         `json:"score"`
    DistinguishType interface{} `json:"distinguishType"`
    IsLocked        bool        `json:"isLocked"`
    IsStickied      bool        `json:"isStickied"`
    Thumbnail       struct {
        URL    string `json:"url"`
        Width  int    `json:"width"`
        Height int    `json:"height"`
    } `json:"thumbnail"`
    Title           string      `json:"title"`
    Author          string      `json:"author"`
    AuthorID        string      `json:"authorId"`
    AuthorIsBlocked bool        `json:"authorIsBlocked"`
    Domain          string      `json:"domain"`
    PostID          string      `json:"postId"`
    UpvoteRatio     float64     `json:"upvoteRatio"`
    NumDuplicates   interface{} `json:"numDuplicates"`
    DiscussionType  interface{} `json:"discussionType"`
    ViewCount       int         `json:"viewCount"`
    GoldCount       int         `json:"goldCount"`
    IsArchived      bool        `json:"isArchived"`
    ContestMode     bool        `json:"contestMode"`
    Gildings        interface{} `json:"gildings"`
    PostCategories  interface{} `json:"postCategories"`
    SuggestedSort   interface{} `json:"suggestedSort"`
    BelongsTo       struct {
        ID   string `json:"id"`
        Type string `json:"type"`
    } `json:"belongsTo"`
    Flair               []interface{} `json:"flair"`
    Hidden              bool          `json:"hidden"`
    Saved               bool          `json:"saved"`
    IsGildable          bool          `json:"isGildable"`
    IsMediaOnly         bool          `json:"isMediaOnly"`
    IsSponsored         bool          `json:"isSponsored"` // 广告过滤掉
    IsNSFW              bool          `json:"isNSFW"`
    IsMeta              bool          `json:"isMeta"`
    IsSpoiler           bool          `json:"isSpoiler"`
    IsBlank             bool          `json:"isBlank"`
    SendReplies         bool          `json:"sendReplies"`
    VoteState           int           `json:"voteState"`
    Permalink           string        `json:"permalink"`
    Events              []interface{} `json:"events"`
    EventsOnRender      []interface{} `json:"eventsOnRender"`
    CallToAction        interface{}   `json:"callToAction"`
    DomainOverride      interface{}   `json:"domainOverride"`
    ImpressionID        interface{}   `json:"impressionId"`
    IsSurveyAd          interface{}   `json:"isSurveyAd"`
    ImpressionIDStr     interface{}   `json:"impressionIdStr"`
    IsCreatedFromAdsUI  bool          `json:"isCreatedFromAdsUi"`
    AdSupplementaryText interface{}   `json:"adSupplementaryText"`
    Subcaption          interface{}   `json:"subcaption"`
    Media               struct {
        Obfuscated  interface{} `json:"obfuscated"`
        Content     string      `json:"content"`
        Type        string      `json:"type"`
        Width       int         `json:"width"`
        Height      int         `json:"height"`
        Resolutions []struct {
            URL    string `json:"url"`
            Width  int    `json:"width"`
            Height int    `json:"height"`
        } `json:"resolutions"`
        RichtextContent struct {
            Document []interface{} `json:"document"`
        } `json:"richtextContent"`
        RteMode string `json:"rteMode"`
    } `json:"media"`
    Preview struct {
        URL    string `json:"url"`
        Width  int    `json:"width"`
        Height int    `json:"height"`
    } `json:"preview"`
    CrosspostRootID       interface{} `json:"crosspostRootId"`
    CrosspostParentID     interface{} `json:"crosspostParentId"`
    NumCrossposts         int         `json:"numCrossposts"`
    IsCrosspostable       bool        `json:"isCrosspostable"`
    LiveCommentsWebsocket string      `json:"liveCommentsWebsocket"`
    Source                interface{} `json:"source"`
    IsOriginalContent     bool        `json:"isOriginalContent"`
    ContentCategories     interface{} `json:"contentCategories"`
    IsScoreHidden         bool        `json:"isScoreHidden"`
}

// RedditIndex 首页数据结构
type RedditIndexData struct {
    Kind string `json:"kind"`
    Data struct {
        After     string      `json:"after"`
        Dist      int         `json:"dist"`
        Modhash   string      `json:"modhash"`
        GeoFilter interface{} `json:"geo_filter"`
        Children  []struct {
            Kind string    `json:"kind"`
            Data IndexData `json:"data"`
        } `json:"children"`
        Before interface{} `json:"before"`
    } `json:"data"`
}

type MediaEmbed struct {
}
type SecureMediaEmbed struct {
}
type Gildings struct {
}

//
type ImageSetting struct {
    URL    string `json:"url"`
    Width  int    `json:"width"`
    Height int    `json:"height"`
}
type Variants struct {
}
type Images struct {
    Source      ImageSetting   `json:"source"`
    Resolutions []ImageSetting `json:"resolutions"`
    Variants    Variants       `json:"variants"`
    ID          string         `json:"id"`
}
type Preview struct {
    Images  []Images `json:"images"`
    Enabled bool     `json:"enabled"`
}
type IndexData struct {
    IsGallery                  bool             `json:"is_gallery"` //  if gallery is disabled
    ApprovedAtUtc              interface{}      `json:"approved_at_utc"`
    Subreddit                  string           `json:"subreddit"`
    Selftext                   string           `json:"selftext"`
    AuthorFullname             string           `json:"author_fullname"`
    Saved                      bool             `json:"saved"`
    ModReasonTitle             interface{}      `json:"mod_reason_title"`
    Gilded                     int              `json:"gilded"`
    Clicked                    bool             `json:"clicked"`
    Title                      string           `json:"title"`
    LinkFlairRichtext          []interface{}    `json:"link_flair_richtext"`
    SubredditNamePrefixed      string           `json:"subreddit_name_prefixed"`
    Hidden                     bool             `json:"hidden"`
    Pwls                       int              `json:"pwls"`
    LinkFlairCSSClass          interface{}      `json:"link_flair_css_class"`
    Downs                      int              `json:"downs"`
    ThumbnailHeight            int              `json:"thumbnail_height"`
    TopAwardedType             interface{}      `json:"top_awarded_type"`
    HideScore                  bool             `json:"hide_score"`
    Name                       string           `json:"name"`
    Quarantine                 bool             `json:"quarantine"`
    LinkFlairTextColor         string           `json:"link_flair_text_color"`
    UpvoteRatio                float64          `json:"upvote_ratio"`
    AuthorFlairBackgroundColor interface{}      `json:"author_flair_background_color"`
    SubredditType              string           `json:"subreddit_type"`
    Ups                        int              `json:"ups"`
    TotalAwardsReceived        int              `json:"total_awards_received"`
    MediaEmbed                 MediaEmbed       `json:"media_embed"`
    ThumbnailWidth             int              `json:"thumbnail_width"`
    AuthorFlairTemplateID      interface{}      `json:"author_flair_template_id"`
    IsOriginalContent          bool             `json:"is_original_content"`
    UserReports                []interface{}    `json:"user_reports"`
    SecureMedia                interface{}      `json:"secure_media"`
    IsRedditMediaDomain        bool             `json:"is_reddit_media_domain"` //非reddit
    IsMeta                     bool             `json:"is_meta"`
    Category                   interface{}      `json:"category"`
    SecureMediaEmbed           SecureMediaEmbed `json:"secure_media_embed"`
    LinkFlairText              interface{}      `json:"link_flair_text"`
    CanModPost                 bool             `json:"can_mod_post"`
    Score                      int              `json:"score"`
    ApprovedBy                 interface{}      `json:"approved_by"`
    IsCreatedFromAdsUI         bool             `json:"is_created_from_ads_ui"`
    AuthorPremium              bool             `json:"author_premium"`
    Thumbnail                  string           `json:"thumbnail"`
    Edited                     bool             `json:"edited"`
    AuthorFlairCSSClass        interface{}      `json:"author_flair_css_class"`
    AuthorFlairRichtext        []interface{}    `json:"author_flair_richtext"`
    Gildings                   Gildings         `json:"gildings"`
    PostHint                   string           `json:"post_hint"`
    ContentCategories          interface{}      `json:"content_categories"`
    IsSelf                     bool             `json:"is_self"`
    ModNote                    interface{}      `json:"mod_note"`
    Created                    int              `json:"created"`
    LinkFlairType              string           `json:"link_flair_type"`
    Wls                        int              `json:"wls"`
    RemovedByCategory          interface{}      `json:"removed_by_category"`
    BannedBy                   interface{}      `json:"banned_by"`
    AuthorFlairType            string           `json:"author_flair_type"`
    Domain                     string           `json:"domain"`
    AllowLiveComments          bool             `json:"allow_live_comments"`
    SelftextHTML               interface{}      `json:"selftext_html"`
    Likes                      interface{}      `json:"likes"`
    SuggestedSort              interface{}      `json:"suggested_sort"`
    BannedAtUtc                interface{}      `json:"banned_at_utc"`
    URLOverriddenByDest        string           `json:"url_overridden_by_dest"`
    ViewCount                  interface{}      `json:"view_count"`
    Archived                   bool             `json:"archived"`
    NoFollow                   bool             `json:"no_follow"`
    IsCrosspostable            bool             `json:"is_crosspostable"`
    Pinned                     bool             `json:"pinned"`
    Over18                     bool             `json:"over_18"`
    Preview                    Preview          `json:"preview"`
    AllAwardings               []interface{}    `json:"all_awardings"`
    Awarders                   []interface{}    `json:"awarders"`
    MediaOnly                  bool             `json:"media_only"`
    CanGild                    bool             `json:"can_gild"`
    Spoiler                    bool             `json:"spoiler"`
    Locked                     bool             `json:"locked"`
    AuthorFlairText            interface{}      `json:"author_flair_text"`
    TreatmentTags              []interface{}    `json:"treatment_tags"`
    Visited                    bool             `json:"visited"`
    RemovedBy                  interface{}      `json:"removed_by"`
    NumReports                 interface{}      `json:"num_reports"`
    Distinguished              interface{}      `json:"distinguished"`
    SubredditID                string           `json:"subreddit_id"`
    AuthorIsBlocked            bool             `json:"author_is_blocked"`
    ModReasonBy                interface{}      `json:"mod_reason_by"`
    RemovalReason              interface{}      `json:"removal_reason"`
    LinkFlairBackgroundColor   string           `json:"link_flair_background_color"`
    ID                         string           `json:"id"`
    IsRobotIndexable           bool             `json:"is_robot_indexable"`
    ReportReasons              interface{}      `json:"report_reasons"`
    Author                     string           `json:"author"`
    DiscussionType             interface{}      `json:"discussion_type"`
    NumComments                int              `json:"num_comments"`
    SendReplies                bool             `json:"send_replies"`
    WhitelistStatus            string           `json:"whitelist_status"`
    ContestMode                bool             `json:"contest_mode"`
    ModReports                 []interface{}    `json:"mod_reports"`
    AuthorPatreonFlair         bool             `json:"author_patreon_flair"`
    AuthorFlairTextColor       interface{}      `json:"author_flair_text_color"`
    Permalink                  string           `json:"permalink"`
    ParentWhitelistStatus      string           `json:"parent_whitelist_status"`
    Stickied                   bool             `json:"stickied"`
    URL                        string           `json:"url"`
    SubredditSubscribers       int              `json:"subreddit_subscribers"`
    CreatedUtc                 int              `json:"created_utc"`
    NumCrossposts              int              `json:"num_crossposts"`
    Media                      interface{}      `json:"media"`
    IsVideo                    bool             `json:"is_video"`
}
