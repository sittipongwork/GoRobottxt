//Author: Sittipong Jungsakul
//Git Repository: https://github.com/sittipongwork
//Created At: 23/11/2016 (DD/MM/YYYY)
//Updated At: 23/11/2016 (DD/MM/YYYY)
package gorobottxt

/*
 * [Robots] ContainData for robots.txt
 * UserAgent    {*Agent}   UserAgent allowed for keep data in robot.txt | default = *
 * Allow        {[]AllowRobot} AllowUrl for robot can be access that
 * Disallow     {[]DisallowRobot} DisllowUrl for robot can't access that
 * Sitemap      {[]SitemapRobot} Add Sitemap of your web
 */
type Robots struct {
	UserAgent *Agent
	Allow     []AllowRobot
	Disallow  []DisallowRobot
	Sitemap   []SitemapRobot
}

/*
 * [Agent] Agent Allowed for keep data your website in robots.txt | All name of agent: (http://www.user-agents.org/)
 */
type Agent struct {
	AgentName string
}

/*
 * [AllowRobot] It's allowed url for robots can keep data in your website
 * Url      {string}    Url bot Allowed
 */
type AllowRobot struct {
	Url string
}

/*
 *[DisallowRobot] It's disallowed url for robots can't keep data in your website
 */
type DisallowRobot struct {
	Url string
}

/*
 *[SitemapRobot] It's sitemap of your website
 */
type SitemapRobot struct {
	Url string
}

func CreateRobot() *Robots {
	rb := &Robots{}
	rb.UserAgent = &Agent{}
	rb.Allow = []AllowRobot{}
	rb.Disallow = []DisallowRobot{}
	return rb
}

func (rb *Robots) AddUserAgent(useragent string) {
	NewUserAgent := &Agent{AgentName: useragent}
	rb.UserAgent = NewUserAgent
}

func (rb *Robots) AddAllowUrl(url string) {
	NewAllowUrl := AllowRobot{Url: url}
	rb.Allow = append(rb.Allow, NewAllowUrl)
}

func (rb *Robots) AddDisallowUrl(url string) {
	NewDissallowUrl := DisallowRobot{Url: url}
	rb.Disallow = append(rb.Disallow, NewDissallowUrl)
}

func (rb *Robots) AddSitemapUrl(url string) {
	SitemapUrl := SitemapRobot{Url: url}
	rb.Sitemap = append(rb.Sitemap, SitemapUrl)
}

type RobotPack struct {
	Robots []Robots
}

func (rb *Robots) RobotGenerateString() string {
	robotstxt := ""
	if len(rb.UserAgent.AgentName) != 0 {
		robotstxt += "User-agent: " + rb.UserAgent.AgentName + "\n"
	} else {
		robotstxt += "You dont have permission who user agent: https://github.com/sittipongwork#adduseragent"
	}
	return robotstxt
}
